package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeConfigMap(t *testing.T) {
	configMap := map[string][]string{
		"name":   {"First", "Last"},
		"salary": {"Wage"},
		"email":  {"Email", "E-mail"},
		"id":     {"Employee-ID"},
	}
	NormalizeMapKeys(configMap)

	assert.Equal(t, 2, len(configMap["name"]))
	assert.Equal(t, "first", configMap["name"][0])
	assert.Equal(t, "last", configMap["name"][1])

	assert.Equal(t, 1, len(configMap["salary"]))
	assert.Equal(t, "wage", configMap["salary"][0])

	assert.Equal(t, 2, len(configMap["email"]))
	assert.Equal(t, "email", configMap["email"][0])
	assert.Equal(t, "e-mail", configMap["email"][1])

	assert.Equal(t, 1, len(configMap["id"]))
	assert.Equal(t, "employee-id", configMap["id"][0])
}
