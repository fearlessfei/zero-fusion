package errx

import (
	"fmt"
)

type CodeError struct {
	errCode uint32
	errMsg  string
}

// GetErrCode 获得错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

// GetErrMsg 获得错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

// WithErrMsg 错误信息
func (e *CodeError) WithErrMsg(errMsg string) *CodeError {
	e.errMsg = errMsg
	return e
}

// Error 错误详情
func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

// NewErrCodeMsg 新建错误
func NewErrCodeMsg(errCode uint32, errMsg string) *CodeError {
	return &CodeError{errCode: errCode, errMsg: errMsg}
}
