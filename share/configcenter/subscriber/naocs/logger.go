package naocs

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/logger"
	"github.com/zeromicro/go-zero/core/logx"
)

type nacosLogger struct{}

var _ logger.Logger = (*nacosLogger)(nil)

func newNacosLogger() logger.Logger {
	return &nacosLogger{}
}

func (n *nacosLogger) Info(args ...interface{}) {
	logx.Info(args...)
}

func (n *nacosLogger) Warn(args ...interface{}) {
	logx.Info(args...)
}

func (n *nacosLogger) Error(args ...interface{}) {
	logx.Error(args...)
}

func (n *nacosLogger) Debug(args ...interface{}) {
	logx.Debug(args...)
}

func (n *nacosLogger) Infof(fmt string, args ...interface{}) {
	logx.Infof(fmt, args...)
}

func (n *nacosLogger) Warnf(fmt string, args ...interface{}) {
	logx.Infof(fmt, args...)
}

func (n *nacosLogger) Errorf(fmt string, args ...interface{}) {
	logx.Errorf(fmt, args...)
}

func (n *nacosLogger) Debugf(fmt string, args ...interface{}) {
	logx.Debugf(fmt, args...)
}
