package common

import (
	"context"

	zerotrace "github.com/zeromicro/go-zero/core/trace"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"

	"zero-fusion/share/trace"
)

type Option func(*options)

type options struct {
	stopAgent bool
}

// StartSpan 开启trace
func StartSpan(ctx context.Context, spanName string, opts ...Option) (context.Context, func(err error)) {
	var o options
	for _, opt := range opts {
		opt(&o)
	}

	tracer := otel.Tracer(trace.TraceName)
	ctx, span := tracer.Start(ctx,
		spanName,
		oteltrace.WithSpanKind(oteltrace.SpanKindClient),
	)

	return ctx, func(err error) {
		if o.stopAgent {
			defer zerotrace.StopAgent()
		}
		defer span.End()

		if err == nil {
			span.SetStatus(codes.Ok, "")
			return
		}

		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
	}
}

// WithStopAgent 停止trace
func WithStopAgent() Option {
	return func(o *options) {
		o.stopAgent = true
	}
}
