package domain

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

const (
	nameMinLen        = 3
	nameMaxLen        = 15
	descriptionMaxLen = 3000
	employeeCntMax    = 10_000_000_000
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

type Company struct {
	ID          uuid.UUID `json:"id" binding:"omitempty"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EmployeeCnt int       `json:"employee_cnt"`
	Registered  bool      `json:"registered"`
	Type        string    `json:"type"`
}

func (c *Company) SetID() {
	c.ID = uuid.New()
}

func (c *Company) Validate() error {
	var errorMsgs []string

	if err := ValidateName(c.Name); err != nil {
		errorMsgs = append(errorMsgs, err.Error())
	}

	if err := ValidateDescription(c.Description); err != nil {
		errorMsgs = append(errorMsgs, err.Error())
	}

	if err := ValidateType(c.Type); err != nil {
		errorMsgs = append(errorMsgs, err.Error())
	}

	if err := ValidateEmployeeCnt(c.EmployeeCnt); err != nil {
		errorMsgs = append(errorMsgs, err.Error())
	}

	if len(errorMsgs) > 0 {
		return fmt.Errorf("%s", strings.Join(errorMsgs, "; "))
	}

	return nil
}

func ValidateEmployeeCnt(cnt int) error {
	if cnt < 0 {
		return fmt.Errorf("amount of employee can't be negative")
	}

	if cnt > employeeCntMax {
		return fmt.Errorf("amount of employee can't exceed %d", employeeCntMax)
	}

	return nil
}

func ValidateName(companyName string) error {
	if len(companyName) < nameMinLen {
		return fmt.Errorf("company name is too short, minimal length is %d", nameMinLen)
	}

	if len(companyName) > nameMaxLen {
		return fmt.Errorf("company name is too long, max length is %d", nameMaxLen)
	}

	return nil
}

func ValidateDescription(companyDescription string) error {
	if len(companyDescription) > descriptionMaxLen {
		return fmt.Errorf("company description is too long, max length is %d", descriptionMaxLen)
	}

	return nil
}

func ValidateType(companyType string) error {
	if availableTypesString == "" {
		availableTypesList := make([]string, 0, len(availableTypes))
		for availableType := range availableTypes {
			availableTypesList = append(availableTypesList, availableType)
		}

		availableTypesString = strings.Join(availableTypesList, ", ")
	}

	if _, ok := availableTypes[companyType]; !ok {
		return fmt.Errorf(
			"\"%s\" company type is not allowed,"+
				" only the following types are allowed: %s",
			companyType,
			availableTypesString,
		)
	}

	return nil
}
