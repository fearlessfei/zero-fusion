package rpcserver

import (
	"context"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"zero-fusion/share/errx"
)

// LoggerInterceptor rpc服务logger
func LoggerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	resp, err = handler(ctx, req)
	if err != nil {
		causeErr := errors.Cause(err)
		var e *errx.CodeError
		if errors.As(causeErr, &e) {
			logx.WithContext(ctx).Errorf("RPC-SRV-ERR %+v", err)
			// 转换到grpc错误
			err = status.Error(codes.Code(e.GetErrCode()), e.GetErrMsg())
		}
	}

	return resp, err
}
