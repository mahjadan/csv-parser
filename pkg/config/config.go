package config

import (
	"employee-csv-parser/pkg/utils"
	"encoding/json"
	"github.com/pkg/errors"
	"io"
	"os"
)

type Loader struct {
	configPath string
}

func NewConfigLoader(configPath string) *Loader {
	return &Loader{
		configPath: configPath,
	}
}

func (c Loader) LoadConfig() (map[string][]string, error) {
	configFile, err := os.Open(c.configPath)
	if err != nil {
		return nil, errors.Wrapf(err, "opening config file")
	}
	defer configFile.Close()

	configMap, err := c.parseConfig(configFile)
	if err != nil {
		return nil, err
	}

	utils.NormalizeMapKeys(configMap)
	return configMap, nil
}

func (c Loader) parseConfig(configFile io.Reader) (map[string][]string, error) {
	var configMap map[string][]string
	decoder := json.NewDecoder(configFile)
	err := decoder.Decode(&configMap)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding config file [%v]", c.configPath)
	}
	return configMap, nil
}
