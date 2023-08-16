package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"rcsv/pkg/utils"
)

type Loader struct {
	configFile         io.Reader
	ValidColumnNames   []string
	InvalidColumnNames []string
	ColumnAliasConfig  map[string][]string
}

func NewConfigLoader(configFile io.Reader, validColumns []string, invalidColumns []string) *Loader {
	return &Loader{
		configFile:         configFile,
		ValidColumnNames:   validColumns,
		InvalidColumnNames: invalidColumns,
	}
}

func (c *Loader) LoadConfig() error {
	configMap, err := c.parseConfig()
	if err != nil {
		return err
	}

	utils.NormalizeMapKeys(configMap)
	if len(configMap) == 0 {
		return errors.New("config file is empty")
	}
	c.ColumnAliasConfig = configMap
	return nil
}

func (c *Loader) parseConfig() (map[string][]string, error) {
	var configMap map[string][]string
	decoder := json.NewDecoder(c.configFile)
	err := decoder.Decode(&configMap)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding config file")
	}
	return configMap, nil
}
