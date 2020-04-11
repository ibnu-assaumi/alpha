package elasticsearch

import (
	"context"
	"net/http"
	"os"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/joho/godotenv"
)

func TestGetESClient(t *testing.T) {

	t.Run("POSITIVE_GET_ES_CLIENT", func(t *testing.T) {
		re := regexp.MustCompile(`^(.*api)`)
		cwd, _ := os.Getwd()
		rootPath := re.Find([]byte(cwd))

		err := godotenv.Load(string(rootPath) + `/.env`)
		assert.NoError(t, err)
		client := GetESClient()
		assert.NotNil(t, client)
		_, code, err := client.Ping(os.Getenv("ELASTIC_HOST_1")).Do(context.Background())
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, code)
	})

	t.Run("PANIC_GET_ES_CLIENT", func(t *testing.T) {
		os.Setenv("ELASTIC_HOST_1", "")
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("The code did not panic")
			}
		}()

		GetESClient()
	})
}
