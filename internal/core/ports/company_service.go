package ports

import (
	"context"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/google/uuid"
)

type CompanyService interface {
	Get(ctx context.Context, id uuid.UUID) (*domain.Company, error)
	Create(ctx context.Context, company *domain.Company) error
	Update(ctx context.Context, id uuid.UUID, fieldsToUpdate map[string]any) error
	Delete(ctx context.Context, id uuid.UUID) error
}
