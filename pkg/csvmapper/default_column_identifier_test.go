package csvmapper_test

import (
	"testing"

	"employee-csv-parser/pkg/csvmapper"
	"github.com/stretchr/testify/assert"
)

func TestDefaultColumnIdentifier_MapColumnToIndexes(t *testing.T) {
	t.Run("AllColumnsMapped", func(t *testing.T) {
		columns := []string{"name", "email", "salary"}
		aliases := map[string][]string{
			"name":   {"Full Name", "Employee Name"},
			"email":  {"E-mail", "Email Address"},
			"salary": {"Wage"},
		}
		identifier := csvmapper.NewDefaultColumnIdentifier()

		err := identifier.MapColumnToIndexes(columns, aliases)

		assert.NoError(t, err)

		t.Run("IndexForColumn", func(t *testing.T) {
			assert.Equal(t, 0, identifier.IndexForColumn("name"))
			assert.Equal(t, 1, identifier.IndexForColumn("email"))
			assert.Equal(t, 2, identifier.IndexForColumn("salary"))
		})
	})
	t.Run("AllColumnsMapped_DifferentOrder", func(t *testing.T) {
		columns := []string{"email", "name", "salary"}
		aliases := map[string][]string{
			"name":   {"Full Name", "Employee Name"},
			"email":  {"E-mail", "Email Address"},
			"salary": {"Wage"},
		}
		identifier := csvmapper.NewDefaultColumnIdentifier()

		err := identifier.MapColumnToIndexes(columns, aliases)

		assert.NoError(t, err)

		t.Run("IndexForColumn", func(t *testing.T) {
			assert.Equal(t, 1, identifier.IndexForColumn("name"))
			assert.Equal(t, 0, identifier.IndexForColumn("email"))
			assert.Equal(t, 2, identifier.IndexForColumn("salary"))
		})
	})

	t.Run("MissingColumns", func(t *testing.T) {
		columns := []string{"name", "email"} // missing salary
		aliases := map[string][]string{
			"name":   {"Full Name", "Employee Name"},
			"email":  {"E-mail", "Email Address"},
			"salary": {"Wage"},
		}
		identifier := csvmapper.NewDefaultColumnIdentifier()

		err := identifier.MapColumnToIndexes(columns, aliases)

		assert.Error(t, err)
		assert.Contains(t, err.Error(), "missing column/header: 'salary'")

		t.Run("IndexForColumn", func(t *testing.T) {
			assert.Equal(t, 0, identifier.IndexForColumn("name"))
			assert.Equal(t, 1, identifier.IndexForColumn("email"))
			assert.Equal(t, -1, identifier.IndexForColumn("salary"))
		})
	})
}
