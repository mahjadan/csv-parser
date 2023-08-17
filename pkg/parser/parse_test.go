package parser

import (
	"bytes"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"rcsv/pkg/config"
	"rcsv/pkg/csvmapper"
	"rcsv/pkg/utils"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	t.Run("ValidInput", func(t *testing.T) {
		csvData := "id,name,email,salary\n1,John,john@example.com,5000\n"
		mockValidFile := &bytes.Buffer{}
		mockInvalidFile := &bytes.Buffer{}
		configLoader := getConfigLoader(t)
		columnIdentifier := csvmapper.NewDefaultColumnIdentifier()

		err := Parse(strings.NewReader(csvData), mockValidFile, mockInvalidFile, configLoader, columnIdentifier)

		assert.NoError(t, err)
		assert.Equal(t, "id,name,email,salary\n1,John,john@example.com,5000\n", mockValidFile.String())
		assert.Equal(t, "id,name,email,salary,errors\n", mockInvalidFile.String())
	})
	t.Run("ValidInputWithDuplicatedRecords", func(t *testing.T) {
		csvData := "id,name,email,salary\n1,John,john@example.com,5000\n2,Mah,mah@test.com,6000\n1,John,john@example.com,5000\n"
		mockValidFile := &bytes.Buffer{}
		mockInvalidFile := &bytes.Buffer{}
		configLoader := getConfigLoader(t)

		columnIdentifier := csvmapper.NewDefaultColumnIdentifier()
		err := Parse(strings.NewReader(csvData), mockValidFile, mockInvalidFile, configLoader, columnIdentifier)

		assert.NoError(t, err)
		assert.Equal(t, "id,name,email,salary\n1,John,john@example.com,5000\n2,Mah,mah@test.com,6000\n", mockValidFile.String())
		assert.Equal(t, "id,name,email,salary,errors\n", mockInvalidFile.String())
	})
	t.Run("ErrorMissingHeaders", func(t *testing.T) {
		csvData := "id,name,email\n1,John,john@example.com\n"
		mockValidFile := &bytes.Buffer{}
		mockInvalidFile := &bytes.Buffer{}
		configLoader := getConfigLoader(t)

		columnIdentifier := csvmapper.NewDefaultColumnIdentifier()

		err := Parse(strings.NewReader(csvData), mockValidFile, mockInvalidFile, configLoader, columnIdentifier)

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "missing column/header: 'salary'")
		assert.Empty(t, mockValidFile.String())
		assert.Empty(t, mockInvalidFile.String())
	})
}

func getConfigLoader(t *testing.T) *config.Loader {
	t.Helper()
	configData := `{
		"column_aliases":{
		  "name": ["Name", "First", "Last", "Full_name", "f.name"],
		  "salary": ["wage", "salary", "pay", "Rate"],
		  "email": ["Email", "E-mail", "e_mail"],
		  "id": ["ID", "emp id","Number"]
		},
		  "output": {
			"valid_columns": ["id", "name", "email", "salary"]
		  }
    }`
	configFile := strings.NewReader(configData)
	v := viper.New()
	v.SetConfigType("json")
	if err := v.ReadConfig(configFile); err != nil {
		t.Fatalf("fail to read config file: %v", err)
	}
	c := config.Config{}
	if err := v.Unmarshal(&c); err != nil {
		t.Fatalf("fail to unmarshal config: %v", err)
	}
	utils.NormalizeMapKeys(c.ColumnAliases)
	return &config.Loader{
		ValidColumnNames:  c.Output.ValidColumns,
		ColumnAliasConfig: c.ColumnAliases,
	}
}
