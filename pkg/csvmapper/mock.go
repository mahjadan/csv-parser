package csvmapper

import "github.com/stretchr/testify/mock"

type MockColumnIdentifier struct {
	mock.Mock
}

func (m *MockColumnIdentifier) IndexForColumn(columnName string) int {
	return m.Called(columnName).Int(0)
}

func (m *MockColumnIdentifier) MapColumnToIndexes(csvHeaders []string, columnAliases map[string][]string) error {
	return m.Called(csvHeaders, columnAliases).Error(0)
}
