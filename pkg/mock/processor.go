package mock

import (
	testifymock "github.com/stretchr/testify/mock"
)

type CSVProcessor struct {
	testifymock.Mock
}

func (m CSVProcessor) ProcessRecord(record []string, err error) bool {
	args := m.Called(record, err)
	return args.Bool(0)
}

func (m CSVProcessor) Flush() {
	m.Called()
}

func (m CSVProcessor) InitializeHeaders(validColumnNames []string, headers []string) error {
	args := m.Called(validColumnNames, headers)
	return args.Error(0)
}

func (m CSVProcessor) GetUniqueRecords() int64 {
	args := m.Called()
	return args.Get(0).(int64)
}
