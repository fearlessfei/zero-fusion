package demo

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"zero-fusion/app/demo/api/scripts/internal/svc"
)

type DemoLogic struct {
	logx.Logger
	Ctx    context.Context
	SvcCtx *svc.ServiceContext
}

func (l *DemoLogic) Run() error {
	l.Logger.Infof("这是一个demo")
	return nil
}
