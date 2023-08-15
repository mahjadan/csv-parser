package processor

import (
	"employee-csv-parser/pkg/csvmapper"
	"employee-csv-parser/pkg/models"
)

type DefaultCSVProcessor struct {
	columnIdentifier csvmapper.ColumnIdentifier
	validWriter      CSVWriter
	invalidWriter    CSVWriter
	processedEmails  map[string]struct{}
}

func NewCSVProcessor(validWriter, invalidWriter CSVWriter, columnIdentifier csvmapper.ColumnIdentifier) *DefaultCSVProcessor {
	return &DefaultCSVProcessor{
		validWriter:      validWriter,
		invalidWriter:    invalidWriter,
		columnIdentifier: columnIdentifier,
		processedEmails:  make(map[string]struct{}),
	}
}

func (p *DefaultCSVProcessor) ProcessValidRecord(record []string) {
	employee := models.Employee{
		ID:     record[p.columnIdentifier.IndexForColumn("id")],
		Email:  record[p.columnIdentifier.IndexForColumn("email")],
		Name:   record[p.columnIdentifier.IndexForColumn("name")],
		Salary: record[p.columnIdentifier.IndexForColumn("salary")],
	}
	err := employee.IsValid()
	if err != nil {
		p.WriteInvalidRecord(record, err.Error())
		return
	}
	if _, exists := p.processedEmails[employee.Email]; !exists {
		p.processedEmails[employee.Email] = struct{}{}
		p.writeValidRecord(employee)
	}
}

func (p *DefaultCSVProcessor) writeValidRecord(employee models.Employee) {
	p.validWriter.Write([]string{employee.ID, employee.Name, employee.Email, employee.Salary})
}

func (p *DefaultCSVProcessor) WriteInvalidRecord(record []string, errorMessage string) {
	p.invalidWriter.Write(append(record, errorMessage))
}

func (p *DefaultCSVProcessor) WriteHeaders(validColumnNames, invalidColumnNames []string) error {
	err := p.validWriter.Write(validColumnNames)
	if err != nil {
		return err
	}
	err = p.invalidWriter.Write(invalidColumnNames)
	if err != nil {
		return err
	}
	return nil
}