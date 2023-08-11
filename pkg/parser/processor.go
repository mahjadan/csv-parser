package parser

//
//import (
//	"employee-csv-parser/pkg/models"
//	"encoding/csv"
//)
//
//type CSVProcessor struct {
//	cfg           map[string][]string
//	columnIndices ColumnIndices
//	validWriter   *csv.Writer
//	invalidWriter *csv.Writer
//}
//
//func NewCSVProcessor(cfg map[string][]string, validWriter, invalidWriter *csv.Writer, columnIndices ColumnIndices) *CSVProcessor {
//	return &CSVProcessor{
//		cfg:           cfg,
//		validWriter:   validWriter,
//		invalidWriter: invalidWriter,
//		columnIndices: columnIndices,
//	}
//}
//func (p *CSVProcessor) ProcessRecord(record []string) error {
//	employee, err := p.constructEmployee(record)
//	if err != nil {
//		p.writeInvalidRecord(record, err.Error())
//		return err
//	}
//
//	err = employee.IsValid()
//	if err != nil {
//		p.writeInvalidRecord(record, err.Error())
//	} else {
//		p.writeValidRecord(employee)
//	}
//
//	return nil
//}
//
//func (p *CSVProcessor) writeValidRecord(employee models.Employee) {
//	p.validWriter.Write([]string{employee.ID, employee.Name, employee.Email, employee.Salary})
//}
//
//func (p *CSVProcessor) writeInvalidRecord(record []string, errorMessage string) {
//	p.invalidWriter.Write(append(record, errorMessage))
//}
//func (p *CSVProcessor) constructEmployee(record []string) models.Employee {
//	return models.Employee{
//		ID:     record[p.columnIndices.GetIndex("id")],
//		Email:  record[p.columnIndices.GetIndex("email")],
//		Name:   record[p.columnIndices.GetIndex("name")],
//		Salary: record[p.columnIndices.GetIndex("salary")],
//	}
//}
