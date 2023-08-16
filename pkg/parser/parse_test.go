package parser

import (
	"bytes"
	"employee-csv-parser/pkg/config"
	"employee-csv-parser/pkg/csvmapper"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		csvData := "id,name,email,salary\n1,John,john@example.com,5000\n"
		mockValidFile := &bytes.Buffer{}
		mockInvalidFile := &bytes.Buffer{}
		configFile := getConfigFile()
		validColumnNames := []string{"id", "name", "email", "salary"}
		invalidColumnNames := append(validColumnNames, "errors")

		columnsConfig := setupConfig(t, configFile, validColumnNames, invalidColumnNames)
		columnIdentifier := csvmapper.NewDefaultColumnIdentifier()

		err := Parse(strings.NewReader(csvData), mockValidFile, mockInvalidFile, columnsConfig, columnIdentifier)

		assert.NoError(t, err)
		assert.Equal(t, "id,name,email,salary\n1,John,john@example.com,5000\n", mockValidFile.String())
		assert.Equal(t, "id,name,email,salary,errors\n", mockInvalidFile.String())
	})
	t.Run("ValidInputWithDuplicatedRecords", func(t *testing.T) {
		csvData := "id,name,email,salary\n1,John,john@example.com,5000\n2,Mah,mah@test.com,6000\n1,John,john@example.com,5000\n"
		mockValidFile := &bytes.Buffer{}
		mockInvalidFile := &bytes.Buffer{}
		configFile := getConfigFile()
		validColumnNames := []string{"id", "name", "email", "salary"}
		invalidColumnNames := append(validColumnNames, "errors")

		columnsConfig := setupConfig(t, configFile, validColumnNames, invalidColumnNames)
		columnIdentifier := csvmapper.NewDefaultColumnIdentifier()

		err := Parse(strings.NewReader(csvData), mockValidFile, mockInvalidFile, columnsConfig, columnIdentifier)

		assert.NoError(t, err)
		assert.Equal(t, "id,name,email,salary\n1,John,john@example.com,5000\n2,Mah,mah@test.com,6000\n", mockValidFile.String())
		assert.Equal(t, "id,name,email,salary,errors\n", mockInvalidFile.String())
	})
	t.Run("ErrorMissingHeaders", func(t *testing.T) {
		csvData := "id,name,email\n1,John,john@example.com\n"
		mockValidFile := &bytes.Buffer{}
		mockInvalidFile := &bytes.Buffer{}
		configFile := getConfigFile()

		validColumnNames := []string{"id", "name", "email", "salary"}
		invalidColumnNames := append(validColumnNames, "errors")
		columnsConfig := setupConfig(t, configFile, validColumnNames, invalidColumnNames)
		columnIdentifier := csvmapper.NewDefaultColumnIdentifier()

		err := Parse(strings.NewReader(csvData), mockValidFile, mockInvalidFile, columnsConfig, columnIdentifier)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "missing column/header: 'salary'")
		assert.Empty(t, mockValidFile.String())
		assert.Empty(t, mockInvalidFile.String())
	})
}

func getConfigFile() *strings.Reader {
	configData := `{
			"name": ["Name", "First", "Last"],
			"salary": ["wage", "salary", "pay"],
			"email": ["Email", "E-mail"],
			"id": ["ID", "emp id"]
		}`
	configFile := strings.NewReader(configData)
	return configFile
}

func setupConfig(t *testing.T, configFile *strings.Reader, validColumnNames []string, invalidColumnNames []string) *config.Loader {
	t.Helper()
	mockColumnsConfig := config.NewConfigLoader(configFile, validColumnNames, invalidColumnNames)
	err := mockColumnsConfig.LoadConfig()
	assert.NoError(t, err)
	return mockColumnsConfig
}
