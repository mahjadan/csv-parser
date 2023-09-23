package mock

import (
	"github.com/stretchr/testify/mock"
)

type Loader struct {
	mock.Mock
}

func (l Loader) GetValidColumnNames() []string {
	args := l.Called()
	return args.Get(0).([]string)
}

func (l Loader) GetColumnAliasMap() map[string][]string {
	args := l.Called()
	return args.Get(0).(map[string][]string)
}
