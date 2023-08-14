package models

import (
	"employee-csv-parser/pkg/csvmapper"
	"github.com/pkg/errors"
	"regexp"
	"strings"
)

const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,4}$`

type Employee struct {
	Name   string
	Email  string
	Salary string
	ID     string
}

func (e Employee) IsValid() error {
	var validationErrors []string

	if e.Name == "" {
		validationErrors = append(validationErrors, "Name is required")
	}
	if e.Email == "" {
		validationErrors = append(validationErrors, "Email is required")
	} else {
		validEmail := regexp.MustCompile(emailRegex)
		if !validEmail.MatchString(e.Email) {
			validationErrors = append(validationErrors, "Invalid email format")
		}
	}
	if e.Salary == "" {
		validationErrors = append(validationErrors, "Salary is required")
	}
	if e.ID == "" {
		validationErrors = append(validationErrors, "ID is required")
	}

	if len(validationErrors) > 0 {
		return errors.New(strings.Join(validationErrors, "| "))
	}
	return nil
}

func NewEmployee(record []string, columnIdentifier csvmapper.ColumnIdentifier) Employee {
	return Employee{
		ID:     record[columnIdentifier.IndexForColumn("id")],
		Email:  record[columnIdentifier.IndexForColumn("email")],
		Name:   record[columnIdentifier.IndexForColumn("name")],
		Salary: record[columnIdentifier.IndexForColumn("salary")],
	}
}
