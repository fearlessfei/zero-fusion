package result

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"

	"zero-fusion/share/errx"
)

// HttpResult 响应结果
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		result := Success(resp)
		httpx.WriteJson(w, http.StatusOK, result)
	} else {
		// 默认错误
		errCode := errx.ServerError.GetErrCode()
		errMsg := "服务器繁忙，请稍后再试"

		causeErr := errors.Cause(err)
		var e *errx.CodeError
		if errors.As(causeErr, &e) {
			errCode = e.GetErrCode()
			errMsg = e.GetErrMsg()
		} else {
			if grpcStatus, ok := status.FromError(causeErr); ok {
				grpcCode := uint32(grpcStatus.Code())

				// grpc错误码详情在grpc codes.go文件中，这里的错误码最小要保证大于codes.go文件中的最大错误码
				// Success.GetErrCode()为业务错误最小值需要确保大于codes.go文件中的最大错误码
				// 其他情况为grpc内部错误
				if grpcCode >= errx.Success.GetErrCode() {
					errCode = grpcCode
					errMsg = grpcStatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("API-SRV-ERR: %+v ", err)

		httpx.WriteJson(w, http.StatusOK, Error(errCode, errMsg))
	}
}

// ParamErrorResult 参数错误
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := errx.RequestParamError.GetErrMsg()
	logx.WithContext(r.Context()).Errorf("%s, %s", errMsg, err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(errx.RequestParamError.GetErrCode(), errMsg))
}
