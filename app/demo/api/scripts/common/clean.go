package common

import (
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/trace"
)

// ExitClean clean exit
func ExitClean(logger logx.Logger) {
	logger.Infof("stopping trace agent")
	trace.StopAgent()
	logger.Infof("trace agent stopped")

	logger.Infof("closing logx")
	err := logx.Close()
	if err != nil {
		logger.Errorf("close logx error: %v", err)
		return
	}
	logger.Infof("logx closed")
}
