package result

import "zero-fusion/share/errx"

type ResponseSuccess struct {
	Code uint32      `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
type NullJson struct{}

func Success(data interface{}) *ResponseSuccess {
	return &ResponseSuccess{errx.Success.GetErrCode(), errx.Success.GetErrMsg(), data}
}

type ResponseError struct {
	Code uint32 `json:"code"`
	Msg  string `json:"msg"`
}

func Error(errCode uint32, errMsg string) *ResponseError {
	return &ResponseError{errCode, errMsg}
}
