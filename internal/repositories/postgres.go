package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"log"
	"time"

	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"

	"github.com/Masterminds/squirrel"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	companyTable = "company"
)

type Postgres struct {
	Pool    *pgxpool.Pool
	Builder squirrel.StatementBuilderType

	dsn          string
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration
}

func New(dsn string, maxPoolSize, connAttempts, connTimeoutSeconds int) (*Postgres, error) {
	pg := Postgres{
		dsn:          dsn,
		maxPoolSize:  maxPoolSize,
		connAttempts: connAttempts,
		connTimeout:  time.Duration(int(time.Second) * connTimeoutSeconds),
	}

	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("error while parsing pool config from dsn: %w", err)
	}
	poolConfig.MaxConns = int32(pg.maxPoolSize)
	for attempts := pg.connAttempts; attempts > 0; attempts-- {
		pg.Pool, err = pgxpool.NewWithConfig(context.Background(), poolConfig)
		if err == nil {
			break
		}
		log.Printf("postgres: trying to connect to database instance, %d attempts left", attempts)
		time.Sleep(pg.connTimeout)
	}

	if err != nil {
		return nil, fmt.Errorf("error while connecting to the database: %w", err)
	}
	log.Printf("postgres: connected successfully")
	return &pg, nil
}

func (p Postgres) Migrate() error {
	var migration *migrate.Migrate
	var err error
	for attempts := p.connAttempts; attempts > 0; attempts-- {
		migration, err = migrate.New("file://migrations", p.dsn)
		if err == nil {
			break
		}
		log.Printf("migrate: trying to connect to database instance, %d attempts left", attempts)
		time.Sleep(p.connTimeout)
	}
	if err != nil {
		return fmt.Errorf("error while trying to migrate: database connection error: %s", err)
	}

	err = migration.Up()
	defer migration.Close()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("migrate: up error: %w", err)
	}

	if errors.Is(err, migrate.ErrNoChange) {
		log.Printf("migrate: no change")
	} else {
		log.Printf("migrate: success")
	}
	return nil
}

func (p Postgres) GetCompanyById(ctx context.Context, id uuid.UUID) (*domain.Company, error) {
	query, args, err := p.Builder.
		Select("id, name, description, employee_cnt, registered, type").
		From(companyTable).
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("get company: error building query: %w", err)
	}

	row := p.Pool.QueryRow(ctx, query, args...)
	targetCompany := new(domain.Company)
	err = row.Scan(
		&targetCompany.Id,
		&targetCompany.Name,
		&targetCompany.Description,
		&targetCompany.EmployeeCnt,
		&targetCompany.Registered,
		&targetCompany.Type,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.NewCompanyNotFoundError(id)
		}
		return nil, fmt.Errorf("get company: error scanning row: %w", err)
	}
	return targetCompany, nil
}

func (p Postgres) CreateCompany(ctx context.Context, company *domain.Company) error {
	if company.Id == uuid.Nil {
		company.SetId()
	}

	query, args, err := p.Builder.
		Insert(companyTable).
		Columns("id, name, description, employee_cnt, registered, type").
		Values(company.Id, company.Name, company.Description, company.EmployeeCnt, company.Registered, company.Type).
		ToSql()
	if err != nil {
		return fmt.Errorf("create company: error building query: %w", err)
	}

	_, err = p.Pool.Exec(ctx, query, args...)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return domain.NewCompanyNameAlreadyTakenError(company.Name)
		}
		return fmt.Errorf("create company: error executing query: %w", err)
	}

	return nil
}

func (p Postgres) UpdateCompany(ctx context.Context, id uuid.UUID, fieldsToUpdate map[string]any) error {
	updateQueryBuilder := p.Builder.Update(companyTable)
	for field, val := range fieldsToUpdate {
		updateQueryBuilder = updateQueryBuilder.Set(field, val)
	}
	updateQueryBuilder = updateQueryBuilder.Where("id = ?", id)
	query, args, err := updateQueryBuilder.ToSql()
	if err != nil {
		return fmt.Errorf("update company: error building query: %w", err)
	}

	commandTag, err := p.Pool.Exec(ctx, query, args...)
	if err != nil {
		var constraintError *pgconn.PgError
		if _, nameChanged := fieldsToUpdate["name"]; nameChanged &&
			errors.As(err, &constraintError) &&
			constraintError.Code == pgerrcode.UniqueViolation {
			return domain.NewCompanyNameAlreadyTakenError(fieldsToUpdate["name"].(string))
		}
		return fmt.Errorf("update company: error executing query: %w", err)
	}
	if commandTag.RowsAffected() == 0 {
		return domain.NewCompanyNotFoundError(id)
	}
	return nil
}

func (p Postgres) DeleteCompany(ctx context.Context, id uuid.UUID) error {
	query, args, err := p.Builder.
		Delete(companyTable).
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("delete company: error building query: %w", err)
	}
	commandTag, err := p.Pool.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("delete company: error executing query: %w", err)
	}

	if commandTag.RowsAffected() == 0 {
		return domain.NewCompanyNotFoundError(id)
	}
	return nil
}
