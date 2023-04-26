package ports

import (
	"context"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/google/uuid"
)

type Repository interface {
	Migrate() error
	GetCompanyById(ctx context.Context, id uuid.UUID) (*domain.Company, error)
	CreateCompany(ctx context.Context, company *domain.Company) error

	// UpdateCompany fetches updated fields and values provided in the fieldsToUpdate argument.
	// Returns either user-friendly errors of types domain.CompanyNotFoundError and
	// domain.NameAlreadyTakenError, or internal error.
	UpdateCompany(ctx context.Context, id uuid.UUID, fieldsToUpdate map[string]any) error
	DeleteCompany(ctx context.Context, id uuid.UUID) error
}
