package config

type Mapper interface {
	GetValidColumnNames() []string
	GetColumnAliasMap() map[string][]string
}
