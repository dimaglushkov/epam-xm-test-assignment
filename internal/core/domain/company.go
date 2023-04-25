package domain

import "github.com/google/uuid"

type Company struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	EmployeeCnt int       `json:"employee_cnt"`
	Registered  bool      `json:"registered"`
	Type        string    `json:"type"`
}
