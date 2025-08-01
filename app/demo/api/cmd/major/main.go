package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/rest"

	"zero-fusion/app/demo/api/internal/config"
	"zero-fusion/app/demo/api/internal/handler"
	"zero-fusion/app/demo/api/internal/svc"
	"zero-fusion/share/swaggerapi"
)

var configFile = flag.String("f", "etc/dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	var runOptions []rest.RunOption
	isDevOrTest := c.Mode == service.DevMode || c.Mode == service.TestMode
	if isDevOrTest {
		// 允许跨域
		runOptions = append(runOptions, rest.WithCors())

		// swagger
		runOptions = swaggerapi.SwaggerRunOptions(runOptions)
	}

	server := rest.MustNewServer(c.RestConf, runOptions...)
	defer server.Stop()

	if isDevOrTest {
		// 日志输出到控制台
		logx.AddWriter(logx.NewWriter(os.Stdout))
	}

	ctx := svc.NewServiceContext(c)
	handler.RegisterMajorHandlers(server, ctx)

	if isDevOrTest {
		// 打印路由
		server.PrintRoutes()
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
