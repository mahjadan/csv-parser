package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalizeMapKeys(t *testing.T) {

	t.Run("NormalizeMapKeys", func(t *testing.T) {
		testCases := []struct {
			name              string
			inputConfigMap    map[string][]string
			expectedConfigMap map[string][]string
		}{
			{
				name: "NormalCase",
				inputConfigMap: map[string][]string{
					"name":   {"First", "Last"},
					"salary": {"Wage"},
					"email":  {"Email", "E-mail"},
					"id":     {"Employee-ID"},
				},
				expectedConfigMap: map[string][]string{
					"name":   {"first", "last"},
					"salary": {"wage"},
					"email":  {"email", "e-mail"},
					"id":     {"employee-id"},
				},
			},
			{
				name:              "EmptyConfigMap",
				inputConfigMap:    map[string][]string{},
				expectedConfigMap: map[string][]string{},
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				NormalizeMapKeys(tc.inputConfigMap)
				assert.Equal(t, tc.expectedConfigMap, tc.inputConfigMap)
			})
		}
	})
}

func TestToLowerTrimSlice(t *testing.T) {
	testCases := []struct {
		name           string
		columnSlice    []string
		expectedResult []string
	}{
		{
			name:           "NormalCase",
			columnSlice:    []string{" First ", " Last ", "\uFEFFEmail"},
			expectedResult: []string{"first", "last", "email"},
		},
		{
			name:           "EmptySlice",
			columnSlice:    []string{},
			expectedResult: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ToLowerTrimSlice(tc.columnSlice)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
