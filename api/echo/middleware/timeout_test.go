package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestTimeoutRequest(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	c := echo.New().NewContext(req, rec)
	t.Run("SUCCESS", func(t *testing.T) {
		handler := func(c echo.Context) error {
			return c.String(200, "")
		}
		handlerFunc := TimeoutRequest(handler)
		handlerFunc(c)
	})
}
