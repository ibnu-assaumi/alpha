package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

func TestHTTPErrorHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	t.Run("ERROR_SYSTEM", func(t *testing.T) {
		errorHandlerFunc := HTTPErrorHandler(e)
		err := &echo.HTTPError{Code: 500, Message: "internal server error"}
		errorHandlerFunc(err, c)
	})

	t.Run("ERROR_HTTP", func(t *testing.T) {
		errorHandlerFunc := HTTPErrorHandler(e)
		err := &echo.HTTPError{Code: 404, Message: "not found"}
		errorHandlerFunc(err, c)
	})
}
