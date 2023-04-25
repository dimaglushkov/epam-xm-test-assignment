package repositories

import (
	"fmt"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(dsn string) (*Postgres, error) {
	db, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to the database: %w", err)
	}

	return &Postgres{
		db: db,
	}, nil
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
