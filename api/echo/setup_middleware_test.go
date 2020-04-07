package echo

import (
	"os"
	"testing"

	"github.com/labstack/echo/v4"
)

func Test_setupMiddleware(t *testing.T) {

	t.Run("POSITIVE_SETUP_MIDDLEWARE_LOG_TRUE", func(t *testing.T) {
		os.Setenv("SERVER_LOG_MODE", "true")
		setupMiddleware(echo.New())
	})

	t.Run("POSITIVE_SETUP_MIDDLEWARE_LOG_FALSE", func(t *testing.T) {
		os.Setenv("SERVER_LOG_MODE", "false")
		setupMiddleware(echo.New())
	})

	t.Run("POSITIVE_SETUP_MIDDLEWARE_LOG_ERRO", func(t *testing.T) {
		os.Setenv("SERVER_LOG_MODE", "")
		setupMiddleware(echo.New())
	})
}
