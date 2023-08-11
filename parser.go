package main

import (
	"employee-csv-parser/pkg/models"
	"encoding/csv"
	"os"
)

func ParseCSV(filePath string) ([]models.Employee, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var employees []models.Employee
	for _, record := range records[1:] {
		employee := models.Employee{
			Name:   record[0],
			Email:  record[1],
			Wage:   record[2],
			Number: record[3],
		}
		employees = append(employees, employee)
	}

	return employees, nil
}
