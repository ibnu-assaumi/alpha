package middleware

import (
	"net/http"
	"runtime/debug"

	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// HTTPErrorHandler : middleware to handle http error
// eg. custom error response on internal server error etc
func HTTPErrorHandler(e *echo.Echo) echo.HTTPErrorHandler {
	return func(err error, c echo.Context) {
		span := opentracing.SpanFromContext(c.Request().Context())
		if span == nil {
			span, _ = opentracing.StartSpanFromContext(c.Request().Context(), "http.error.handler")
		}
		defer span.Finish()
		span.SetTag("error.message", err.Error())

		code := http.StatusInternalServerError
		if he, ok := err.(*echo.HTTPError); ok {
			code = he.Code
		}

		ext.HTTPStatusCode.Set(span, uint16(code))
		ext.Error.Set(span, true)
		if code == http.StatusInternalServerError {
			span.SetTag("error.stacktrace", string(debug.Stack()))
		}
		c.Logger().Error(err)
		e.DefaultHTTPErrorHandler(err, c)
	}
}
