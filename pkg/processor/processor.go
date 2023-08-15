package processor

import (
	"employee-csv-parser/pkg/csvmapper"
	"employee-csv-parser/pkg/models"
	"encoding/csv"
)

type DefaultCSVRecordProcessor struct {
	columnIdentifier csvmapper.ColumnIdentifier
	validWriter      *csv.Writer
	invalidWriter    *csv.Writer
	processedEmails  map[string]struct{}
}

func NewCSVProcessor(validWriter, invalidWriter *csv.Writer, columnIdentifier csvmapper.ColumnIdentifier) *DefaultCSVRecordProcessor {
	return &DefaultCSVRecordProcessor{
		validWriter:      validWriter,
		invalidWriter:    invalidWriter,
		columnIdentifier: columnIdentifier,
		processedEmails:  make(map[string]struct{}),
	}
}
func (p *DefaultCSVRecordProcessor) ProcessValidRecord(record []string) {
	employee := models.NewEmployee(record, p.columnIdentifier)
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

func (p *DefaultCSVRecordProcessor) writeValidRecord(employee models.Employee) {
	p.validWriter.Write([]string{employee.ID, employee.Name, employee.Email, employee.Salary})
}

func (p *DefaultCSVRecordProcessor) WriteInvalidRecord(record []string, errorMessage string) {
	p.invalidWriter.Write(append(record, errorMessage))
}

func (p *DefaultCSVRecordProcessor) WriteHeaders(validColumnNames, invalidColumnNames []string) error {
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

type CSVProcessor interface {
	ProcessValidRecord(record []string) error
	InitializeHeaders() error
}
