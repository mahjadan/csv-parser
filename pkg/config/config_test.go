package config

import (
	"github.com/stretchr/testify/assert"
	"io"
	"strings"
	"testing"
)

func TestNewConfigLoader(t *testing.T) {
	testCases := []struct {
		name           string
		configFile     io.Reader
		validColumns   []string
		invalidColumns []string
		expected       *Loader
	}{
		{
			name:           "WithValidColumns",
			configFile:     strings.NewReader("{}"),
			validColumns:   []string{"name", "email"},
			invalidColumns: []string{"salary", "id"},
			expected: &Loader{
				configFile:         strings.NewReader("{}"),
				ValidColumnNames:   []string{"name", "email"},
				InvalidColumnNames: []string{"salary", "id"},
			},
		},
		{
			name:           "WithoutColumns",
			configFile:     nil,
			validColumns:   nil,
			invalidColumns: nil,
			expected: &Loader{
				configFile:         nil,
				ValidColumnNames:   nil,
				InvalidColumnNames: nil,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			configLoader := NewConfigLoader(tc.configFile, tc.validColumns, tc.invalidColumns)
			assert.Equal(t, tc.expected, configLoader)
		})
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("ParseExistingConfigFile", func(t *testing.T) {
		configData := `{
			"name": ["Name", "First", "Last"],
			"salary": ["wage", "salary", "pay"],
			"email": ["Email", "E-mail"],
			"id": ["ID", "emp id"]
		}`
		configFile := strings.NewReader(configData)

		configLoader := NewConfigLoader(configFile, nil, nil)
		err := configLoader.LoadConfig()

		assert.NoError(t, err)
		assert.NotNil(t, configLoader.ColumnAliasConfig)
		assert.Equal(t, 3, len(configLoader.ColumnAliasConfig["name"]))
		assert.Equal(t, 3, len(configLoader.ColumnAliasConfig["salary"]))
		assert.Equal(t, 2, len(configLoader.ColumnAliasConfig["email"]))
		assert.Equal(t, 2, len(configLoader.ColumnAliasConfig["id"]))
	})
	t.Run("InvalidJSONConfig", func(t *testing.T) {
		tempConfig := `{
			"name": inavlid,
			"id": "ID"
		}`
		configFile := strings.NewReader(tempConfig)

		configLoader := NewConfigLoader(configFile, nil, nil)
		err := configLoader.LoadConfig()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error decoding config file")
		assert.Nil(t, configLoader.ColumnAliasConfig)
	})
	t.Run("ValidEmptyJSONConfig", func(t *testing.T) {
		tempConfig := `{}`
		configFile := strings.NewReader(tempConfig)

		configLoader := NewConfigLoader(configFile, nil, nil)
		err := configLoader.LoadConfig()

		assert.NotNil(t, err)
		assert.Contains(t, err.Error(), "config file is empty")
		assert.Nil(t, configLoader.ColumnAliasConfig)
	})
}
