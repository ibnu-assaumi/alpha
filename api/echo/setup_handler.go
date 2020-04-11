package echo

import (
	"github.com/labstack/echo/v4"

	handlerBeta "github.com/Bhinneka/alpha/api/echo/handler/v1/beta"
	handlerCharlie "github.com/Bhinneka/alpha/api/echo/handler/v1/charlie"
)

func setupHandler(e *echo.Echo) {
	// group handler charlie v1
	v1GroupCharlie := e.Group("/v1/charlie")
	handlerCharlie.NewHandler().Mount(v1GroupCharlie)
	// group handler beta v1
	v1GroupBeta := e.Group("/v1/beta")
	handlerBeta.NewHandler().Mount(v1GroupBeta)
}
