package parser

import (
	"employee-csv-parser/pkg/models"
	"employee-csv-parser/pkg/utils"
	"encoding/csv"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"time"
)

var standardOutputColumns = []string{"id", "name", "email", "salary"}

func Parse(cfg map[string][]string, csfFilePath string) error {
	csvFile, err := os.Open(csfFilePath)
	if err != nil {
		return errors.Wrap(err, "error opening CSV file")
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV headers:", err)
		return errors.Wrap(err, "error reading CSV headers")
	}
	headers = utils.ToLowerTrimSlice(headers)

	columnIndices := NewColumnIndices(cfg, standardOutputColumns)
	err = columnIndices.MapIndices(headers)
	if err != nil {
		return err
	}
	validFileName := generateFileName("valid")
	invalidFileName := generateFileName("invalid")

	validFile, err := os.Create(validFileName)
	if err != nil {
		return errors.Wrap(err, "error creating valid CSV file")
	}
	defer validFile.Close()

	invalidFile, err := os.Create(invalidFileName)
	if err != nil {
		return errors.Wrap(err, "error creating invalid CSV file")
	}
	defer invalidFile.Close()

	// create writers
	validWriter := csv.NewWriter(validFile)
	invalidWriter := csv.NewWriter(invalidFile)
	defer validWriter.Flush()
	defer invalidWriter.Flush()

	// Write headers to the valid CSV file
	validWriter.Write(standardOutputColumns)
	invalidHeaders := append(headers, "errors")
	invalidWriter.Write(invalidHeaders)

	var employee models.Employee
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			// handle wrong number of fields
			record = append(record, err.Error())
			invalidWriter.Write(record)
			continue
		}

		employee = models.Employee{}

		// Map columns based on headers and columnMappings
		employee.ID = record[columnIndices.GetIndex("id")]
		employee.Email = record[columnIndices.GetIndex("email")]
		employee.Name = record[columnIndices.GetIndex("name")]
		employee.Salary = record[columnIndices.GetIndex("salary")]
		err = employee.IsValid()
		if err != nil {
			fmt.Printf("INVALID: Processed Employee: %+v, error: %v \n", employee, err)
			invalidWriter.Write(append(record, err.Error()))
		} else {
			fmt.Printf("Processed Employee: %+v\n", employee)
			validWriter.Write([]string{employee.ID, employee.Name, employee.Email, employee.Salary})
		}
	}
	return nil
}
func generateFileName(prefix string) string {
	timestamp := time.Now().Format("20060102150405")
	return fmt.Sprintf("%s_%s.csv", prefix, timestamp)
}

type ColumnIndices struct {
	indices        map[string]int
	columnMappings map[string][]string
	columnNames    []string
}

func (c *ColumnIndices) GetIndex(columnName string) int {
	return c.indices[columnName]
}
func NewColumnIndices(cfg map[string][]string, columnNames []string) ColumnIndices {
	indices := make(map[string]int)
	for columnName := range cfg {
		indices[columnName] = -1
	}
	return ColumnIndices{indices: indices, columnMappings: cfg, columnNames: columnNames}
}
func (c *ColumnIndices) MapIndices(csvHeaders []string) error {
	for index, header := range csvHeaders {
		for columnName, alternativeNames := range c.columnMappings {
			if utils.SliceContains(alternativeNames, header) {
				c.indices[columnName] = index
			}
		}
	}
	return c.hasMissingColumn()
}

func (c *ColumnIndices) hasMissingColumn() error {
	for columnName, index := range c.indices {
		if index == -1 {
			return errors.Errorf("missing column/header: '%s'. Tried alternatives: %v", columnName, c.columnMappings[columnName])
		}
	}
	return nil
}
