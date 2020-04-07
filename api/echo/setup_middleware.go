package echo

import (
	"os"
	"strconv"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"

	"github.com/Bhinneka/alpha/api/echo/middleware"
)

// setup middleware
func setupMiddleware(e *echo.Echo) {
	_, err := strconv.ParseBool(os.Getenv("SERVER_LOG_MODE"))
	if err == nil {
		e.Use(echoMiddleware.Logger())
	}
	e.Use(echoMiddleware.Recover())
	e.Use(middleware.TimeoutRequest)
	e.Use(echo.WrapMiddleware(middleware.Jaeger))
	e.Use(echoMiddleware.BodyDump(middleware.BodyDump))
	e.HTTPErrorHandler = middleware.HTTPErrorHandler(e)
	e.Use(sentryecho.New(sentryecho.Options{
		Repanic: true,
	}))
}
