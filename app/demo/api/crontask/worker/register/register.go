package register

import (
	"zero-fusion/app/demo/api/crontask/worker/internal/svc"
	"zero-fusion/share/cron/asynqx/worker"
)

// RegisterTask 注册任务
func RegisterTask(ws *worker.WorkerServer, svcCtx *svc.ServiceContext) {
	registerOnceTask(ws, svcCtx)
	registerPeriodicTask(ws, svcCtx)
}
