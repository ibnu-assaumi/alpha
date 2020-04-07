package middleware

import (
	"errors"
	"math"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

type errReader int

func (errReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("test error")
}

func initJaegaer() {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
		},
		ServiceName: "test",
	}
	t, _, _ := cfg.NewTracer(config.MaxTagValueLength(math.MaxInt32))
	opentracing.SetGlobalTracer(t)
}

func TestJaeger(t *testing.T) {

	t.Run("SUCCESS_EXTRACT_HEADER", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		initJaegaer()
		span, _ := opentracing.StartSpanFromContext(req.Context(), "test")
		opentracing.Tracer.Inject(span.Tracer(), span.Context(), opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Header))

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
		})
		h := Jaeger(handler)
		h.ServeHTTP(rec, req)
	})

	t.Run("NEGATIVE_EXTRACT_HEADER", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/", nil)
		rec := httptest.NewRecorder()

		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "application/json")
		})
		h := Jaeger(handler)
		h.ServeHTTP(rec, req)
	})
}
