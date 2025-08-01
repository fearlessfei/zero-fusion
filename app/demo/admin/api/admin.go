package main

import (
	"flag"
	"fmt"
	"zero-fusion/app/demo/admin/api/internal/config"
	"zero-fusion/app/demo/admin/api/internal/handler"
	"zero-fusion/app/demo/admin/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "api/etc/admin.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	server.Use(ctx.AddLog)

	handler.RegisterHandlers(server, ctx)

	// 自定义错误
	//httpx.SetErrorHandler(func(err error) (int, interface{}) {
	//	switch e := err.(type) {
	//	case *errorx.CodeError:
	//		return http.StatusOK, e.Data()
	//	default:
	//		fmt.Println(err.Error())
	//		return http.StatusInternalServerError, nil
	//	}
	//})

	server.PrintRoutes() //print registered routes in rest servers
	logx.DisableStat()
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
