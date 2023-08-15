package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

func TestNewConfigLoader(t *testing.T) {
	testCases := []struct {
		name           string
		configPath     string
		validColumns   []string
		invalidColumns []string
		expected       *Loader
	}{
		{
			name:           "WithValidColumns",
			configPath:     "/path/to/config.json",
			validColumns:   []string{"name", "email"},
			invalidColumns: []string{"salary", "id"},
			expected: &Loader{
				configPath:         "/path/to/config.json",
				ValidColumnNames:   []string{"name", "email"},
				InvalidColumnNames: []string{"salary", "id"},
			},
		},
		{
			name:           "WithoutColumns",
			configPath:     "/path/to/config.json",
			validColumns:   nil,
			invalidColumns: nil,
			expected: &Loader{
				configPath:         "/path/to/config.json",
				ValidColumnNames:   nil,
				InvalidColumnNames: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			configLoader := NewConfigLoader(tc.configPath, tc.validColumns, tc.invalidColumns)
			assert.Equal(t, tc.expected, configLoader)
		})
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("ExistingConfigFile", func(t *testing.T) {

		tempConfig := []byte(`{
		"name": ["Name", "First", "Last"],
		"salary": ["wage", "salary", "pay"],
		"email": ["Email", "E-mail"],
		"id": ["ID", "emp id"]
	}`)

		dir := t.TempDir()
		tempFile, err := os.CreateTemp(dir, "config_test*.json")
		if err != nil {
			t.Fatalf("error creating temp file: %v\n", err)
		}
		defer tempFile.Close()

		_, err = tempFile.Write(tempConfig)
		if err != nil {
			t.Fatal(err)
		}

		configPath := tempFile.Name()
		configLoader := NewConfigLoader(configPath, nil, nil)
		err = configLoader.LoadConfig()

		assert.NoError(t, err)
		assert.NotNil(t, configLoader.ColumnAliasConfig)
		assert.Equal(t, 3, len(configLoader.ColumnAliasConfig["name"]))
		assert.Equal(t, 3, len(configLoader.ColumnAliasConfig["salary"]))
		assert.Equal(t, 2, len(configLoader.ColumnAliasConfig["email"]))
		assert.Equal(t, 2, len(configLoader.ColumnAliasConfig["id"]))
	})
	t.Run("NonExistentConfigFile", func(t *testing.T) {
		configPath := "/path/to/nonexistent/config.json"
		configLoader := NewConfigLoader(configPath, nil, nil)
		err := configLoader.LoadConfig()

		assert.Error(t, err)
		assert.Nil(t, configLoader.ColumnAliasConfig)
		assert.Contains(t, err.Error(), "opening config file")
	})
	t.Run("InvalidJSONConfig", func(t *testing.T) {
		tempConfig := []byte(`{
		"name": inavlid,
		"id": "ID"
	}`)

		dir := t.TempDir()
		tempFile, err := os.CreateTemp(dir, "config_test*.json")
		if err != nil {
			t.Fatalf("error creating temp file: %v\n", err)
		}
		defer tempFile.Close()

		_, err = tempFile.Write(tempConfig)
		if err != nil {
			t.Fatal(err)
		}

		configPath := tempFile.Name()
		configLoader := NewConfigLoader(configPath, nil, nil)
		err = configLoader.LoadConfig()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error decoding config file")
		assert.Nil(t, configLoader.ColumnAliasConfig)
	})
}

func TestParseConfig(t *testing.T) {
	configData := `{
 		"name":   ["First", "Last"],
		"salary": ["Wage"],
		"email":  ["Email", "E-mail"],
		"id":     ["Employee-ID"]
	}`
	reader := strings.NewReader(configData)
	configLoader := &Loader{}
	parseConfig, err := configLoader.parseConfig(reader)

	assert.NoError(t, err)
	assert.NotNil(t, parseConfig)
	assert.Equal(t, 2, len(parseConfig["name"]))
	assert.Equal(t, 1, len(parseConfig["salary"]))

	assert.Equal(t, 2, len(parseConfig["email"]))
	assert.Equal(t, 1, len(parseConfig["id"]))
}
