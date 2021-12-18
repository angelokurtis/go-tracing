package tracing

import (
	"github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/config"
)

type SamplerType string

const (
	ConstantSampler      SamplerType = jaeger.SamplerTypeConst
	ProbabilisticSampler SamplerType = jaeger.SamplerTypeProbabilistic
	RateLimitingSampler  SamplerType = jaeger.SamplerTypeRateLimiting
	RemoteSampler        SamplerType = jaeger.SamplerTypeRemote
)

// JAEGER_SAMPLER_TYPE=const
// JAEGER_SAMPLER_PARAM="1"
// JAEGER_ENDPOINT=http://jaeger-collector.lvh.me/api/traces

type TracerOptionFunc func(*config.Configuration)

func WithServiceName(service string) TracerOptionFunc {
	return func(o *config.Configuration) {
		o.ServiceName = service
	}
}

func WithSamplerType(samplerType SamplerType) TracerOptionFunc {
	return func(o *config.Configuration) {
		o.Sampler = samplerType
	}
}

func WithSamplerParam(samplerParam float32) TracerOptionFunc {
	return func(o *config.Configuration) {
		o.SamplerParam = samplerParam
	}
}

func WithEndpoint(endpoint string) TracerOptionFunc {
	return func(o *config.Configuration) {
		o.Endpoint = endpoint
	}
}
