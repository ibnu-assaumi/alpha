package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func Test_handler_Check(t *testing.T) {
	t.Run("POSITIVE_HEALTH_CHECK", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()
		c := echo.New().NewContext(req, rec)

		h := NewHandler()
		if assert.NoError(t, h.Check(c)) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
}
