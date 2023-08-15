package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"strings"
	"testing"
)

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
		defer func() {
			err := os.Remove(tempFile.Name())
			if err != nil {
				t.Fatal(err)
			}
		}()
		defer tempFile.Close()

		_, err = tempFile.Write(tempConfig)
		if err != nil {
			t.Fatal(err)
		}

		configPath := tempFile.Name()
		configLoader := NewConfigLoader(configPath)
		config, err := configLoader.LoadConfig()

		assert.NoError(t, err)
		assert.NotNil(t, config)
		assert.Equal(t, 3, len(config["name"]))
		assert.Equal(t, 3, len(config["salary"]))
		assert.Equal(t, 2, len(config["email"]))
		assert.Equal(t, 2, len(config["id"]))
	})
	t.Run("NonExistentConfigFile", func(t *testing.T) {
		configPath := "/path/to/nonexistent/config.json"
		configLoader := NewConfigLoader(configPath)
		configMap, err := configLoader.LoadConfig()

		assert.Error(t, err)
		assert.Nil(t, configMap)
		assert.Contains(t, err.Error(), "opening config file")
	})
	t.Run("InvalidJSONConfig", func(t *testing.T) {
		tempConfig := []byte(`{
		invalid JSON content
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
		configLoader := NewConfigLoader(configPath)
		config, err := configLoader.LoadConfig()

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error decoding config file")
		assert.Nil(t, config)
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
