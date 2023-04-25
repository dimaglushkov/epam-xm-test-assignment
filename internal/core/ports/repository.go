package ports

import (
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/google/uuid"
)

type RepositoryPort interface {
	Migrate() error
	GetCompanyById(id uuid.UUID) (*domain.Company, error)
	CreateCompany(company domain.Company) error
	UpdateCompany(company domain.Company) error
	DeleteCompany(company domain.Company) error
}
