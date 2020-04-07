package jsonschema

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	type testSchema struct {
		ID        int    `json:"id"`
		StringVal string `json:"stringVal"`
	}

	t.Run("NO_ERROR", func(t *testing.T) {
		param := testSchema{
			ID:        1,
			StringVal: "value",
		}
		err := Validate("schema_test.json", param)
		assert.Nil(t, err)
	})

	t.Run("ERROR_OPEN_FILE", func(t *testing.T) {
		err := Validate("schema_test", testSchema{})
		assert.NotNil(t, err)
	})

	t.Run("INVALID", func(t *testing.T) {
		err := Validate("schema_test.json", testSchema{})
		fmt.Println(err)
		assert.NotNil(t, err)
	})
}
