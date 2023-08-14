package config

import (
	"employee-csv-parser/pkg/utils"
	"encoding/json"
	"github.com/pkg/errors"
	"os"
)

type Loader struct {
	configPath string
	config     map[string][]string
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

	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&c.config)
	if err != nil {
		return nil, errors.Wrapf(err, "error decoding config file [%v]", c.configPath)
	}

	c.normalizeConfigMap()
	return c.config, nil
}
func (c Loader) normalizeConfigMap() {
	for k, v := range c.config {
		c.config[k] = utils.ToLowerTrimSlice(v)
	}
}
