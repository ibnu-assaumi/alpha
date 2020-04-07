package charlie

import (
	"testing"

	"github.com/labstack/echo/v4"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	t.Run("POSITIVE_NEW_HANDLER", func(t *testing.T) {
		assert.NotNil(t, NewHandler())
	})
}

func Test_handler_Mount(t *testing.T) {
	h := NewHandler()
	assert.NotNil(t, h)

	group := echo.New().Group("test")
	h.Mount(group)
}
