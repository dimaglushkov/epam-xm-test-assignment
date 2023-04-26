package services

import (
	"context"
	"errors"
	"fmt"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/ports"
	"github.com/google/uuid"
	"log"
)

type CompanyService struct {
	repo ports.RepositoryPort
	ps   ports.PubSubPort
}

func NewCompanyService(repo ports.RepositoryPort, ps ports.PubSubPort) *CompanyService {
	return &CompanyService{
		repo: repo,
		ps:   ps,
	}
}

func (cs CompanyService) Get(ctx context.Context, id uuid.UUID) (*domain.Company, error) {
	company, err := cs.repo.GetCompanyById(ctx, id)
	if err != nil {
		var companyNotFoundErr *domain.CompanyNotFoundError
		if errors.As(err, &companyNotFoundErr) {
			return nil, err
		}
		log.Printf("company service: get: %s", err.Error())
		return nil, domain.ErrInternalServer
	}
	return company, err
}

func (cs CompanyService) Create(ctx context.Context, company *domain.Company) error {
	if err := company.Validate(); err != nil {
		return fmt.Errorf("validation error: %w", err)
	}

	company.SetId()
	err := cs.repo.CreateCompany(ctx, company)
	if err != nil {
		var companyNameAlreadyTakenError *domain.CompanyNameAlreadyTakenError
		if errors.As(err, &companyNameAlreadyTakenError) {
			return err
		}
		log.Printf("company service: create: %s", err.Error())
		return domain.ErrInternalServer
	}

	return nil
}

// Update when using gin.Context's Bind method it's impossible to distinguish
// skipped fields from fields intentionally set to zero value, thus
// updating requires manual validation of the received fields
func (cs CompanyService) Update(ctx context.Context, id uuid.UUID, fieldsToUpdate map[string]any) error {
	for field, val := range fieldsToUpdate {
		switch field {
		case "name":
			if _, ok := val.(string); !ok {
				return fmt.Errorf("unsupported type for field %s", field)
			}
			if err := domain.ValidateName(val.(string)); err != nil {
				return err
			}

		case "description":
			if _, ok := val.(string); !ok {
				return fmt.Errorf("unsupported type for field %s", field)
			}
			if err := domain.ValidateDescription(val.(string)); err != nil {
				return err
			}

		case "type":
			if _, ok := val.(string); !ok {
				return fmt.Errorf("unsupported type for field %s", field)
			}
			if err := domain.ValidateType(val.(string)); err != nil {
				return err
			}

		case "employee_cnt":
			_, okInt := val.(int)
			_, okFloat := val.(float64)
			if !okInt && !okFloat {
				return fmt.Errorf("unsupported type for field %s", field)
			}
			if okFloat {
				fieldsToUpdate[field] = int(val.(float64))
			}

		case "registered":
			if _, ok := val.(bool); !ok {
				return fmt.Errorf("unsupported type for field %s", field)
			}

		default:
			delete(fieldsToUpdate, field)
		}
	}

	err := cs.repo.UpdateCompany(ctx, id, fieldsToUpdate)
	if err != nil {
		var companyNameAlreadyTakenError *domain.CompanyNameAlreadyTakenError
		var companyNotFoundErr *domain.CompanyNotFoundError
		if errors.As(err, &companyNotFoundErr) || errors.As(err, &companyNameAlreadyTakenError) {
			return err
		}
		log.Printf("company service: update: %s", err.Error())
		return domain.ErrInternalServer
	}
	return nil
}

func (cs CompanyService) Delete(ctx context.Context, id uuid.UUID) error {
	err := cs.repo.DeleteCompany(ctx, id)
	if err != nil {
		var companyNotFoundErr *domain.CompanyNotFoundError

		if errors.As(err, &companyNotFoundErr) {
			return err
		}
		log.Printf("company service: delete: %s", err.Error())
		return domain.ErrInternalServer
	}
	return nil
}
