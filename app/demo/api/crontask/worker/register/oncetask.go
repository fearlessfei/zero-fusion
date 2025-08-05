package register

import (
	"zero-fusion/app/demo/api/crontask/worker/internal/logic/oncetaskdemo"
	"zero-fusion/app/demo/api/crontask/worker/internal/svc"
	"zero-fusion/share/cron/asynqx/worker"
)

// registerOnceTask 注册一次性任务
func registerOnceTask(ws *worker.WorkerServer, svcCtx *svc.ServiceContext) {
	ws.RegisterTask(oncetaskdemo.NewDemoLogic(svcCtx))
}
