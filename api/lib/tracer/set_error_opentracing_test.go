package tracer

import (
	"testing"

	"github.com/opentracing/opentracing-go"
)

func TestSetErrorOpentracing(t *testing.T) {
	t.Run("POSITIVE_SET_ERROR_OPENTRACING", func(t *testing.T) {
		span := opentracing.StartSpan("test")
		SetErrorOpentracing(span, "test", "test")
	})
}
