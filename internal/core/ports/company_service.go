package ports

import (
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/google/uuid"
)

type CompanyServicePort interface {
	GetById(id uuid.UUID) (*domain.Company, error)
	Create(company domain.Company) error
	Update(company domain.Company) error
	Delete(company domain.Company) error
}
