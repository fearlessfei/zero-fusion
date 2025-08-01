package log

import (
	"os"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
)

// OutStdout 输出日志到标准输出
func OutStdout(mode string) {
	isDevOrTest := mode == service.DevMode || mode == service.TestMode
	if isDevOrTest {
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}
}
