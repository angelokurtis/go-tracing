package tracing

import (
	"fmt"
	"io"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
)

// Initialize create an instance of Jaeger Tracer and sets it as GlobalTracer.
func Initialize(module string, options ...TracerOptionFunc) (io.Closer, error) {
	set(module, module)
	cfg := new(config.Configuration)

	for _, fn := range options {
		if fn == nil {
			continue
		}
		fn(cfg)
	}

	cfg, err := cfg.FromEnv()
	if err != nil {
		return nil, fmt.Errorf("cannot read environment variables: %w", err)
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, fmt.Errorf("cannot create tracer: %w", err)
	}

	opentracing.SetGlobalTracer(tracer)
	return closer, nil
}
