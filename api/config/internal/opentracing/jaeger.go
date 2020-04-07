package opentracing

import (
	"fmt"
	"math"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

// InitOpenTracing : init global open tracing
func InitOpenTracing() {
	cfg := &config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort: fmt.Sprintf("%s:%s",
				os.Getenv("JAEGER_AGENT_HOST"),
				os.Getenv("JAEGER_AGENT_PORT"),
			),
		},
		ServiceName: os.Getenv("JAEGER_SERVICE_NAME"),
	}
	t, _, err := cfg.NewTracer(config.MaxTagValueLength(math.MaxInt32))
	if err != nil {
		fmt.Printf("ERROR: cannot init opentracing connection: %v\n", err)
		panic(err)
	}
	opentracing.SetGlobalTracer(t)
}
