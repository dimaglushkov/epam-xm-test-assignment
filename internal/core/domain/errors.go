package domain

import (
	"fmt"

	"github.com/google/uuid"
)

var ErrInternalServer = fmt.Errorf("internal server error")

type CompanyNotFoundError struct {
	ID uuid.UUID
}

func NewCompanyNotFoundError(id uuid.UUID) *CompanyNotFoundError {
	return &CompanyNotFoundError{ID: id}
}

func (e CompanyNotFoundError) Error() string {
	return fmt.Sprintf("company with id \"%s\" doesnt exist", e.ID)
}

type NameAlreadyTakenError struct {
	Name string
}

func NewNameAlreadyTakenError(name string) *NameAlreadyTakenError {
	return &NameAlreadyTakenError{Name: name}
}

func (e NameAlreadyTakenError) Error() string {
	return fmt.Sprintf("company with the name \"%s\" already exists", e.Name)
}
