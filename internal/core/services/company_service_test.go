package services_test

import (
	"context"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/domain"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/core/services"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/events"
	"github.com/dimaglushkov/epam-xm-test-assignment/internal/repositories"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"testing"
)

var (
	appName = "test-app"
	Company = domain.Company{
		Id:          uuid.New(),
		Name:        "some name",
		Description: "",
		EmployeeCnt: 6,
		Registered:  true,
		Type:        "NonProfit",
	}
	ids = [...]uuid.UUID{uuid.New(), uuid.New(), uuid.New()}

	CompanyNotFoundError         *domain.CompanyNotFoundError
	CompanyNameAlreadyTakenError *domain.NameAlreadyTakenError
)

func TestCompanyService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventsWriter := events.NewMockEventsWriter(ctrl)
	mockRepo := repositories.NewMockRepository(ctrl)
	mockRepo.EXPECT().GetCompanyById(context.Background(), Company.Id).Return(&Company, nil)
	for _, id := range ids {
		mockRepo.EXPECT().GetCompanyById(context.Background(), id).Return(nil, domain.NewCompanyNotFoundError(id))
	}

	companyService := services.NewCompanyService(appName, mockRepo, mockEventsWriter)

	for _, id := range ids {
		res, err := companyService.Get(context.Background(), id)
		assert.Nil(t, res)
		assert.ErrorAs(t, err, &CompanyNotFoundError)
	}

	res, err := companyService.Get(context.Background(), Company.Id)
	assert.Equal(t, res, &Company)
	assert.NoError(t, err)
}

func TestCompanyService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockEventsWriter := events.NewMockEventsWriter(ctrl)
	mockEventsWriter.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil)

	mockRepo := repositories.NewMockRepository(ctrl)
	mockRepo.EXPECT().CreateCompany(gomock.Any(), &Company).Return(nil)
	mockRepo.EXPECT().CreateCompany(gomock.Any(), &Company).Return(domain.NewNameAlreadyTakenError(Company.Name))
	mockRepo.EXPECT().CreateCompany(gomock.Any(), &Company).Return(domain.NewNameAlreadyTakenError(Company.Name))

	companyService := services.NewCompanyService(appName, mockRepo, mockEventsWriter)

	err := companyService.Create(context.Background(), &Company)
	assert.NoError(t, err)
	err = companyService.Create(context.Background(), &Company)
	assert.ErrorAs(t, err, &CompanyNameAlreadyTakenError)
	err = companyService.Create(context.Background(), &Company)
	assert.ErrorAs(t, err, &CompanyNameAlreadyTakenError)
}

func TestCompanyService_Update(t *testing.T) {
	var (
		Update = map[string]any{
			"name":         "new name",
			"random field": 123,
			"registered":   true,
		}
		InvalidUpdate = map[string]any{
			"name":       123,
			"registered": "no",
		}
	)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventsWriter := events.NewMockEventsWriter(ctrl)
	mockEventsWriter.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockRepo := repositories.NewMockRepository(ctrl)
	mockRepo.EXPECT().UpdateCompany(gomock.Any(), Company.Id, Update).Return(nil)
	mockRepo.EXPECT().UpdateCompany(gomock.Any(), ids[0], Update).Return(domain.NewCompanyNotFoundError(ids[0]))
	mockRepo.EXPECT().UpdateCompany(gomock.Any(), ids[1], Update).Return(domain.NewNameAlreadyTakenError(Update["name"].(string)))

	companyService := services.NewCompanyService(appName, mockRepo, mockEventsWriter)

	err := companyService.Update(context.Background(), Company.Id, Update)
	assert.NoError(t, err)
	err = companyService.Update(context.Background(), ids[0], Update)
	assert.ErrorAs(t, err, &CompanyNotFoundError)
	err = companyService.Update(context.Background(), ids[1], Update)
	assert.ErrorAs(t, err, &CompanyNameAlreadyTakenError)
	err = companyService.Update(context.Background(), Company.Id, InvalidUpdate)
	assert.ErrorContains(t, err, "unsupported type")
}

func TestCompanyService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventsWriter := events.NewMockEventsWriter(ctrl)
	mockEventsWriter.EXPECT().Write(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	mockRepo := repositories.NewMockRepository(ctrl)
	mockRepo.EXPECT().DeleteCompany(gomock.Any(), Company.Id).Return(nil)
	mockRepo.EXPECT().DeleteCompany(gomock.Any(), Company.Id).Return(domain.NewCompanyNotFoundError(Company.Id))

	companyService := services.NewCompanyService(appName, mockRepo, mockEventsWriter)
	err := companyService.Delete(context.Background(), Company.Id)
	assert.NoError(t, err)
	err = companyService.Delete(context.Background(), Company.Id)
	assert.ErrorAs(t, err, &CompanyNotFoundError)
}
