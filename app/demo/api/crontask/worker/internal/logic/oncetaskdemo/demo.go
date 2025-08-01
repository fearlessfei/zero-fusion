package oncetaskdemo

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"zero-fusion/app/demo/api/crontask"
	"zero-fusion/app/demo/api/crontask/worker/internal/svc"
	"zero-fusion/share/cron"
)

type DemoLogic struct {
	SvcCtx *svc.ServiceContext
}

func NewDemoLogic(svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		SvcCtx: svcCtx,
	}
}

func (l *DemoLogic) Name() string {
	return crontask.TaskTypeOnceDemo
}

func (l *DemoLogic) Run(ctx context.Context, payload []byte, _ cron.TaskContext) error {
	logger := logx.WithContext(ctx)
	logger.Infof("once payload: %s", payload)
	return nil
}
