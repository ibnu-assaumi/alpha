package echo

import (
	"testing"

	"github.com/labstack/echo/v4"
)

func Test_setupHandler(t *testing.T) {
	t.Run("POSITIVE_SETUP_HANDLER", func(t *testing.T) {
		setupHandler(echo.New())
	})
}
