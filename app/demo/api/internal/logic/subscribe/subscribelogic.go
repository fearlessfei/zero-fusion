package subscribe

import (
	"context"

	"zero-fusion/app/demo/api/internal/svc"
	"zero-fusion/app/demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubscribeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 订阅消息
func NewSubscribeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubscribeLogic {
	return &SubscribeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubscribeLogic) Subscribe(req *types.SubscribeMessageReq) (resp *types.SubscribeMessageResp, err error) {
	// todo: add your logic here and delete this line

	return
}
