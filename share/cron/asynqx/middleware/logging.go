package middleware

import (
	"context"
	"time"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func LoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		logger := logx.WithContext(ctx)
		logger.Infof("Start processing [%s]", t.Type())

		start := time.Now()
		err := h.ProcessTask(ctx, t)
		if err != nil {
			return err
		}

		logger.Infof("Finished processing [%s]: Elapsed Time = %v", t.Type(), time.Since(start))
		return nil
	})
}
