package worker

import (
	"context"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"

	"zero-fusion/share/cron"
	"zero-fusion/share/cron/asynqx/middleware"
)

type WorkerServer struct {
	srv *asynq.Server
	mux *asynq.ServeMux
}

func (w *WorkerServer) RegisterTask(task cron.Task) {
	taskCopy := task
	w.mux.HandleFunc(taskCopy.Name(), func(ctx context.Context, t *asynq.Task) error {
		return taskCopy.Run(ctx, t.Payload(), nil)
	})
	logx.Infof("registered task: %s", task.Name())
}

func (w *WorkerServer) Start() {
	if err := w.srv.Run(w.mux); err != nil {
		panic(err)
	}
}

func (w *WorkerServer) Stop() {
	w.srv.Shutdown()
	logx.Info("worker stopped")
}

func NewWorkerServer(redisClientOpt asynq.RedisClientOpt, config asynq.Config) *WorkerServer {
	srv := asynq.NewServer(
		redisClientOpt,
		config,
	)

	mux := asynq.NewServeMux()
	mux.Use(middleware.TraceMiddleware)
	mux.Use(middleware.LoggingMiddleware)

	return &WorkerServer{
		srv: srv,
		mux: mux,
	}
}
