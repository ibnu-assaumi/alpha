package elasticsearch

import (
	"os"
	"testing"

	"github.com/Bhinneka/alpha/api/config/internal/internaltest"
	"github.com/stretchr/testify/assert"
)

func TestGetESClient(t *testing.T) {
	t.Parallel()
	t.Run("POSITIVE_GET_ES_CLIENT", func(t *testing.T) {
		internaltest.LoadEnv()
		assert.NotNil(t, GetESClient())
	})

	t.Run("PANIC_GET_ES_CLIENT", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		os.Setenv("ELASTIC_7_HOST", "1.2.3.4:abcd")

		GetESClient()
	})

	t.Run("PANIC_GET_ES_CLIENT_INFO", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		os.Setenv("ELASTIC_7_HOST", "localhost:3210")

		GetESClient()
	})
}
