package main

import (
	"employee-csv-parser/pkg/models"
	"encoding/csv"
	"fmt"
	"os"
)

func PrintSummary(validEmployees, invalidEmployees []models.Employee) {
	fmt.Println("Valid Employees:")
	for _, emp := range validEmployees {
		fmt.Printf("Name: %s, Email: %s, Wage: %s, Number: %s\n", emp.Name, emp.Email, emp.Wage, emp.Number)
	}

	fmt.Println("\nInvalid Employees:")
	for _, emp := range invalidEmployees {
		fmt.Printf("Name: %s, Email: %s, Wage: %s, Number: %s\n", emp.Name, emp.Email, emp.Wage, emp.Number)
	}
}

func GenerateCSV(validEmployees, invalidEmployees []models.Employee) error {
	validFile, err := os.Create("valid_employees.csv")
	if err != nil {
		return err
	}
	defer validFile.Close()

	invalidFile, err := os.Create("invalid_employees.csv")
	if err != nil {
		return err
	}
	defer invalidFile.Close()

	validWriter := csv.NewWriter(validFile)
	defer validWriter.Flush()

	invalidWriter := csv.NewWriter(invalidFile)
	defer invalidWriter.Flush()

	for _, emp := range validEmployees {
		validWriter.Write([]string{emp.Name, emp.Email, emp.Wage, emp.Number})
	}

	for _, emp := range invalidEmployees {
		invalidWriter.Write([]string{emp.Name, emp.Email, emp.Wage, emp.Number})
	}

	return nil
}
