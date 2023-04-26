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

type NameAlreadyTakenError struct {
	Name string
}

func NewNameAlreadyTakenError(name string) *NameAlreadyTakenError {
	return &NameAlreadyTakenError{Name: name}
}

func (e NameAlreadyTakenError) Error() string {
	return fmt.Sprintf("company with the name \"%s\" already exists", e.Name)
}
