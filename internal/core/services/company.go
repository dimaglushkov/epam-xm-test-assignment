package services

import (
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/ports"
	"github.com/google/uuid"
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

func (s CompanyService) GetById(id uuid.UUID) (*domain.Company, error) {
	//TODO implement me
	panic("implement me")
}

func (s CompanyService) Create(company domain.Company) error {
	//TODO implement me
	panic("implement me")
}

func (s CompanyService) Update(company domain.Company) error {
	//TODO implement me
	panic("implement me")
}

func (s CompanyService) Delete(company domain.Company) error {
	//TODO implement me
	panic("implement me")
}
