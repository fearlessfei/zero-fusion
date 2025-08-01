package middleware

import (
	"context"
	"zero-fusion/share/trace"

	"github.com/hibiken/asynq"
)

func TraceMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		spanCtx, endSpan := trace.StartSpan(ctx, "asynq-worker")

		err := h.ProcessTask(spanCtx, t)
		endSpan(err)
		if err != nil {
			return err
		}
		return nil
	})
}
