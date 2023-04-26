package domain

import (
	"fmt"
	"github.com/google/uuid"
)

var (
	ErrInternalServer = fmt.Errorf("internal server error")
)

type CompanyNotFoundError struct {
	Id uuid.UUID
}

func NewCompanyNotFoundError(id uuid.UUID) *CompanyNotFoundError {
	return &CompanyNotFoundError{Id: id}
}

func (e CompanyNotFoundError) Error() string {
	return fmt.Sprintf("company with id \"%s\" doesnt exist", e.Id)
}

type CompanyNameAlreadyTakenError struct {
	Name string
}

func NewCompanyNameAlreadyTakenError(name string) *CompanyNameAlreadyTakenError {
	return &CompanyNameAlreadyTakenError{Name: name}
}

func (e CompanyNameAlreadyTakenError) Error() string {
	return fmt.Sprintf("company with the name \"%s\" already exists", e.Name)
}
