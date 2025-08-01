package trace

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	oteltrace "go.opentelemetry.io/otel/trace"
)

// StartSpan 开启trace
func StartSpan(ctx context.Context, spanName string) (context.Context, func(err error)) {
	tracer := otel.Tracer(TraceName)
	ctx, span := tracer.Start(ctx,
		spanName,
		oteltrace.WithSpanKind(oteltrace.SpanKindClient),
	)

	return ctx, func(err error) {
		defer span.End()

		if err == nil {
			span.SetStatus(codes.Ok, "")
			return
		}

		span.SetStatus(codes.Error, err.Error())
		span.RecordError(err)
	}
}
