package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"io"
	"rcsv/pkg/utils"
)

type csvLoader struct {
	configFile        io.Reader
	ValidColumnNames  []string
	ColumnAliasConfig map[string][]string
}

func (l csvLoader) GetValidColumnNames() []string {
	return l.ValidColumnNames
}

func (l csvLoader) GetColumnAliasMap() map[string][]string {
	return l.ColumnAliasConfig
}

func NewConfigLoader(configPath string) (Mapper, error) {
	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("json")
	if err := v.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "error reading config file")
	}

	var config struct {
		ColumnAliases map[string][]string `mapstructure:"column_aliases"`
		Output        struct {
			ValidColumns []string `mapstructure:"valid_columns"`
		} `mapstructure:"output"`
	}
	if err := v.Unmarshal(&config); err != nil {
		return nil, errors.Wrap(err, "Error unmarshalling config")
	}
	utils.NormalizeMapKeys(config.ColumnAliases)
	return &csvLoader{
		ValidColumnNames:  config.Output.ValidColumns,
		ColumnAliasConfig: config.ColumnAliases,
	}, nil
}
