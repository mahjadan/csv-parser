package parser_test

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	mck "rcsv/pkg/mock"
	"rcsv/pkg/parser"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		inputData := strings.NewReader("id,name,email,salary\n1,John,john@example.com,5000\n")
		configLoaderMock := new(mck.Loader)
		columnIdentifierMock := new(mck.MockColumnIdentifier)
		csvProcessorMock := new(mck.CSVProcessor)

		csvProcessorMock.On("InitializeHeaders", []string{"id", "name", "email", "salary"}, []string{"id", "name", "email", "salary", "errors"}).Return(nil)
		csvProcessorMock.On("Flush")
		csvProcessorMock.On("GetUniqueRecords").Return(int64(0))
		csvProcessorMock.On("ProcessRecord", []string{"1", "John", "john@example.com", "5000"}, nil).Return(true).Once()

		columnIdentifierMock.On("MapColumnToIndexes", []string{"id", "name", "email", "salary"}, map[string][]string{}).Return(nil)

		configLoaderMock.On("GetValidColumnNames").Return([]string{"id", "name", "email", "salary"})
		configLoaderMock.On("GetColumnAliasMap").Return(map[string][]string{})

		_, err := parser.Parse(inputData, configLoaderMock, columnIdentifierMock, csvProcessorMock)
		assert.NoError(t, err)
		configLoaderMock.AssertExpectations(t)
		columnIdentifierMock.AssertExpectations(t)
		csvProcessorMock.AssertExpectations(t)
	})
	t.Run("error when initializing headers", func(t *testing.T) {
		inputData := strings.NewReader("id,name,email,salary\n1,John,john@example.com,5000\n")
		configLoaderMock := new(mck.Loader)
		columnIdentifierMock := new(mck.MockColumnIdentifier)
		csvProcessorMock := new(mck.CSVProcessor)

		csvProcessorMock.On("InitializeHeaders", []string{"id", "name", "email", "salary"}, []string{"id", "name", "email", "salary", "errors"}).Return(errors.New("error while writing to file"))
		csvProcessorMock.On("Flush")
		csvProcessorMock.On("GetUniqueRecords").Return(int64(0))
		csvProcessorMock.On("ProcessRecord", []string{"1", "John", "john@example.com", "5000"}, nil).Return(true).Once()

		columnIdentifierMock.On("MapColumnToIndexes", []string{"id", "name", "email", "salary"}, map[string][]string{}).Return(nil)

		configLoaderMock.On("GetValidColumnNames").Return([]string{"id", "name", "email", "salary"})
		configLoaderMock.On("GetColumnAliasMap").Return(map[string][]string{})

		_, err := parser.Parse(inputData, configLoaderMock, columnIdentifierMock, csvProcessorMock)
		assert.Error(t, err)
		configLoaderMock.AssertExpectations(t)
		columnIdentifierMock.AssertExpectations(t)
		csvProcessorMock.AssertNotCalled(t, "ProcessRecord", []string{"1", "John", "john@example.com", "5000"}, nil)
	})
	t.Run("ValidInputWithDuplicatedRecords", func(t *testing.T) {
		data := "id,name,email,salary\n1,John,john@example.com,5000\n2,Mah,mah@test.com,6000\n1,John,john@example.com,5000\n"
		inputData := strings.NewReader(data)
		configLoaderMock := new(mck.Loader)
		columnIdentifierMock := new(mck.MockColumnIdentifier)
		csvProcessorMock := new(mck.CSVProcessor)

		csvProcessorMock.On("InitializeHeaders", []string{"id", "name", "email", "salary"}, []string{"id", "name", "email", "salary", "errors"}).Return(nil)
		csvProcessorMock.On("Flush")
		csvProcessorMock.On("GetUniqueRecords").Return(int64(0))
		csvProcessorMock.On("ProcessRecord", []string{"1", "John", "john@example.com", "5000"}, nil).Return(true).Twice()
		csvProcessorMock.On("ProcessRecord", []string{"2", "Mah", "mah@test.com", "6000"}, nil).Return(true).Once()

		columnIdentifierMock.On("MapColumnToIndexes", []string{"id", "name", "email", "salary"}, map[string][]string{}).Return(nil)

		configLoaderMock.On("GetValidColumnNames").Return([]string{"id", "name", "email", "salary"})
		configLoaderMock.On("GetColumnAliasMap").Return(map[string][]string{})

		stats, err := parser.Parse(inputData, configLoaderMock, columnIdentifierMock, csvProcessorMock)
		assert.NoError(t, err)
		configLoaderMock.AssertExpectations(t)
		columnIdentifierMock.AssertExpectations(t)
		csvProcessorMock.AssertExpectations(t)
		assert.Equal(t, int64(3), stats.TotalValidRecords)

	})
	t.Run("WithSomeInvalidInput", func(t *testing.T) {
		data := "id,name,email,salary\n1,John,john@example.com,5000\n2,Mah,mah@test.com,6000\n1,Jay,jay@example.com,\n"
		inputData := strings.NewReader(data)
		configLoaderMock := new(mck.Loader)
		columnIdentifierMock := new(mck.MockColumnIdentifier)
		csvProcessorMock := new(mck.CSVProcessor)

		csvProcessorMock.On("InitializeHeaders", []string{"id", "name", "email", "salary"}, []string{"id", "name", "email", "salary", "errors"}).Return(nil)
		csvProcessorMock.On("Flush")
		csvProcessorMock.On("GetUniqueRecords").Return(int64(0))
		csvProcessorMock.On("ProcessRecord", []string{"1", "John", "john@example.com", "5000"}, nil).Return(true).Once()
		csvProcessorMock.On("ProcessRecord", []string{"1", "Jay", "jay@example.com", ""}, nil).Return(false).Once()
		csvProcessorMock.On("ProcessRecord", []string{"2", "Mah", "mah@test.com", "6000"}, nil).Return(true).Once()

		columnIdentifierMock.On("MapColumnToIndexes", []string{"id", "name", "email", "salary"}, map[string][]string{}).Return(nil)

		configLoaderMock.On("GetValidColumnNames").Return([]string{"id", "name", "email", "salary"})
		configLoaderMock.On("GetColumnAliasMap").Return(map[string][]string{})

		stats, err := parser.Parse(inputData, configLoaderMock, columnIdentifierMock, csvProcessorMock)
		assert.NoError(t, err)
		configLoaderMock.AssertExpectations(t)
		columnIdentifierMock.AssertExpectations(t)
		csvProcessorMock.AssertExpectations(t)
		assert.Equal(t, int64(2), stats.TotalValidRecords)
		assert.Equal(t, int64(1), stats.TotalInvalidRecords)

	})
	t.Run("error mapping column to index", func(t *testing.T) {
		inputData := strings.NewReader("id,name,salary\n1,John,john@example.com,5000\n")
		configLoaderMock := new(mck.Loader)
		columnIdentifierMock := new(mck.MockColumnIdentifier)
		csvProcessorMock := new(mck.CSVProcessor)

		csvProcessorMock.On("InitializeHeaders", []string{"id", "name", "salary"}, []string{"id", "name", "email", "salary", "errors"}).Return(nil)
		csvProcessorMock.On("Flush")
		csvProcessorMock.On("GetUniqueRecords").Return(int64(0))
		csvProcessorMock.On("ProcessRecord", []string{"1", "John", "john@example.com", "5000"}, nil).Return(true).Once()

		columnIdentifierMock.On("MapColumnToIndexes", []string{"id", "name", "salary"}, map[string][]string{}).Return(errors.New("missing column/header: email"))

		configLoaderMock.On("GetValidColumnNames").Return([]string{"id", "name", "email", "salary"})
		configLoaderMock.On("GetColumnAliasMap").Return(map[string][]string{})

		_, err := parser.Parse(inputData, configLoaderMock, columnIdentifierMock, csvProcessorMock)
		assert.Error(t, err)
		columnIdentifierMock.AssertExpectations(t)
		configLoaderMock.AssertNotCalled(t, "GetValidColumnNames")
		csvProcessorMock.AssertNotCalled(t, "ProcessRecord", []string{"1", "John", "john@example.com", "5000"}, nil)
	})
}
