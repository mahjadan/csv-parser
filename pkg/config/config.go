package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"rcsv/pkg/utils"
)

type Loader struct {
	configFile        io.Reader
	ValidColumnNames  []string
	ColumnAliasConfig map[string][]string
}

type Config struct {
	ColumnAliases map[string][]string `mapstructure:"column_aliases"`
	Output        struct {
		ValidColumns []string `mapstructure:"valid_columns"`
	} `mapstructure:"output"`
}

func NewConfigLoader(configPath string) (*Loader, error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "error reading config file")
	}

	var config Config
	if err := v.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling config")
	}
	utils.NormalizeMapKeys(config.ColumnAliases)
	return &Loader{
		ValidColumnNames:  config.Output.ValidColumns,
		ColumnAliasConfig: config.ColumnAliases,
	}, nil
}
