package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
)

// BodyDump : dump body for tracer to trace request
func BodyDump(c echo.Context, req []byte, res []byte) {
	span := opentracing.SpanFromContext(c.Request().Context())
	if span == nil {
		span, _ = opentracing.StartSpanFromContext(c.Request().Context(), "body.dump")
	}
	defer span.Finish()
	span.SetTag("request.ip.address", c.RealIP())
	span.SetTag("request.referer", c.Request().Referer())
	span.SetTag("response.body", string(res))

}
