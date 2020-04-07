package tracer

import (
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// SetErrorOpentracing : set error opentracing
func SetErrorOpentracing(span opentracing.Span, errorType string, errorMessage interface{}) {
	if span != nil {
		span.SetTag("error.type", errorType)
		span.SetTag("error.message", errorMessage)
		ext.Error.Set(span, true)
	}
}
