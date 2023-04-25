package repositories

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type Postgres struct {
	Pool *pgxpool.Pool
	dsn  string

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

func (p Postgres) GetCompanyById(id uuid.UUID) (*domain.Company, error) {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) CreateCompany(company domain.Company) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) DeleteCompany(company domain.Company) error {
	//TODO implement me
	panic("implement me")
}

func (p Postgres) UpdateCompany(company domain.Company) error {
	//TODO implement me
	panic("implement me")
}
