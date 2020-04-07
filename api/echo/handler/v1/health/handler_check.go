package health

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func (impl handler) Check(c echo.Context) error {
	return c.JSON(http.StatusOK, time.Now().String())
}
