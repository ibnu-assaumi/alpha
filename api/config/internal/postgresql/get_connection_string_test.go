package postgresql

import (
	"testing"

	"github.com/Bhinneka/alpha/api/config/internal/internaltest"
	"github.com/stretchr/testify/assert"
)

func Test_getConnectionString(t *testing.T) {
	t.Parallel()
	t.Run("POSITIVE_GET_CONNECTION_STING", func(t *testing.T) {
		internaltest.LoadEnv()
		connectionString := getConnectionString(dbTypeRead)
		assert.NotNil(t, connectionString)
		assert.NotEqual(t, "", connectionString)
	})

	t.Run("PANIC_GET_CONNECTION_STRING", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		getConnectionString("")
	})
}
