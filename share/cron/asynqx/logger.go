package asynqx

import "github.com/zeromicro/go-zero/core/logx"

type AsynqLogger struct{}

func (l *AsynqLogger) Debug(args ...interface{}) {
	logx.Debug(args...)
}

func (l *AsynqLogger) Info(args ...interface{}) {
	logx.Info(args...)
}

func (l *AsynqLogger) Warn(args ...interface{}) {
	logx.Info(args...)
}

func (l *AsynqLogger) Error(args ...interface{}) {
	logx.Error(args...)
}

func (l *AsynqLogger) Fatal(args ...interface{}) {
	logx.Error(args...)
}
