package domain

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
)

const (
	nameMinLen        = 3
	nameMaxLen        = 15
	descriptionMaxLen = 3000
)

var (
	availableTypes = map[string]any{
		"Corporations":        struct{}{},
		"NonProfit":           struct{}{},
		"Cooperative":         struct{}{},
		"Sole Proprietorship": struct{}{},
	}
	availableTypesString string
)

func init() {
	availableTypesList := make([]string, 0, len(availableTypes))
	for availableType := range availableTypes {
		availableTypesList = append(availableTypesList, availableType)
	}
	availableTypesString = strings.Join(availableTypesList, ", ")
}

type Company struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	EmployeeCnt int       `json:"employee_cnt" binding:"required"`
	Registered  bool      `json:"registered" binding:"required"`
	Type        string    `json:"type" binding:"required"`
}

func (c Company) Validate() []error {
	var errors []error
	if len(c.Name) < nameMinLen {
		errors = append(errors, fmt.Errorf("company name is too short, minimal length is %d", nameMinLen))
	}
	if len(c.Name) > nameMaxLen {
		errors = append(errors, fmt.Errorf("company name is too long, max length is %d", nameMaxLen))
	}
	if len(c.Description) > descriptionMaxLen {
		errors = append(errors, fmt.Errorf("company description is too long, max length is %d", descriptionMaxLen))
	}
	if _, ok := availableTypes[c.Type]; !ok {
		errors = append(errors, fmt.Errorf(
			"\"%s\" company type is not allowed,"+
				" only the following types are allowed: %s",
			c.Type,
			availableTypesString,
		))
	}

	return errors
}
