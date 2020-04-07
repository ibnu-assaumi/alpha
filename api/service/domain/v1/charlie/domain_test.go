package charlie

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDomain_TableName(t *testing.T) {
	t.Run("POSITIVE_TABLE_NAME", func(t *testing.T) {
		obj := &Domain{}
		assert.Equal(t, TableName, obj.TableName())
	})
}

func TestDomain_MarshalJSON(t *testing.T) {
	t.Parallel()
	t.Run("POSITIVE_MARSHAL_JSON", func(t *testing.T) {
		obj := &Domain{}
		data, err := obj.MarshalJSON()
		assert.NoError(t, err)

		result, err := json.Marshal(obj)
		assert.NoError(t, err)

		assert.Equal(t, result, data)
	})
}
