package register

import (
	"zero-fusion/app/demo/api/crontask/worker/internal/logic/periodictaskdemo"
	"zero-fusion/app/demo/api/crontask/worker/internal/svc"
	"zero-fusion/share/cron/asynqx/worker"
)

// registerPeriodicTask 注册周期性定时任务
func registerPeriodicTask(ws *worker.WorkerServer, svcCtx *svc.ServiceContext) {
	ws.RegisterTask(periodictaskdemo.NewDemoLogic(svcCtx))
}
