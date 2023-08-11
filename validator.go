package main

import (
	"employee-csv-parser/pkg/models"
	"net/mail"
)

func ValidateEmail(email string) bool {
	// Implement email validation logic using regex
	// Example: return regexp.MatchString(pattern, email)
	_, err := mail.ParseAddress(email)
	return err == nil
}

func ProcessData(employees []models.Employee) ([]models.Employee, []models.Employee) {
	var validEmployees []models.Employee
	var invalidEmployees []models.Employee

	for _, emp := range employees {
		if ValidateEmail(emp.Email) {
			validEmployees = append(validEmployees, emp)
		} else {
			invalidEmployees = append(invalidEmployees, emp)
		}
	}

	return validEmployees, invalidEmployees
}
