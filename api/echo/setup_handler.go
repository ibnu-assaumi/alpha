package echo

import (
	"github.com/labstack/echo/v4"

	handlerCharlie "github.com/Bhinneka/alpha/api/echo/handler/v1/charlie"
)

func setupHandler(e *echo.Echo) {
	// group handler charlie v1
	v1GroupCharlie := e.Group("/v1/charlie")
	handlerCharlie.NewHandler().Mount(v1GroupCharlie)
}
