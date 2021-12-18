package tracing

import (
	"context"
	"fmt"
	"net/http"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func StartSpanFromContext(ctx context.Context, options ...SpanOptionFunc) (*Span, context.Context) {
	opt := new(SpanOptions)

	for _, fn := range options {
		if fn == nil {
			continue
		}
		fn(opt)
	}

	span, ctx := opentracing.StartSpanFromContext(ctx, opt.Operation())
	if opt.resource != nil {
		span.SetTag("kubernetes.resource", opt.resource.String())
	}
	return &Span{Span: span}, ctx
}

func StartSpanFromRequest(r *http.Request) (*Span, context.Context) {
	ctx := ExtractSpanContextFromRequest(r)
	span, ctxWithSpan := opentracing.StartSpanFromContext(r.Context(), r.Method+" "+r.URL.Path, ext.RPCServerOption(ctx))

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	ext.HTTPUrl.Set(span, fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI))
	ext.HTTPMethod.Set(span, r.Method)
	span.SetTag("http.protocol", r.Proto)

	return &Span{Span: span}, ctxWithSpan
}

func SpanFromContext(ctx context.Context) *Span {
	span := opentracing.SpanFromContext(ctx)
	if span == nil {
		return nil
	}

	return &Span{Span: span}
}

func ExtractSpanContextFromRequest(r *http.Request) opentracing.SpanContext {
	tracer := opentracing.GlobalTracer()
	ctx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	return ctx
}
