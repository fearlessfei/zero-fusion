package main

import (
	"flag"
	"time"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/conf"

	"zero-fusion/app/demo/api/crontask/scheduler/internal/config"
	"zero-fusion/app/demo/api/crontask/scheduler/internal/logic/providerrepo"
	"zero-fusion/app/demo/api/crontask/scheduler/internal/svc"
	"zero-fusion/share/cron/asynqx"
	"zero-fusion/share/cron/asynqx/scheduler"
	"zero-fusion/share/cron/asynqx/scheduler/provider"
	"zero-fusion/share/log"
)

var configFile = flag.String("f", "etc/dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.MustSetUp()
	log.OutStdout(c.Mode)

	// 时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	ds := scheduler.NewShedulerDynamicServer(
		asynq.RedisClientOpt{
			Addr:     c.AsynqRedis.Host,
			Password: c.AsynqRedis.Pass,
		},
		&asynq.SchedulerOpts{
			Logger:          &asynqx.AsynqLogger{},
			Location:        loc,
			PreEnqueueFunc:  scheduler.PreEnqueueFunc,
			PostEnqueueFunc: scheduler.PostEnqueueFunc,
		},
		&provider.ConfigProvider{
			Repo: &providerrepo.DBRepo{SvcCtx: svc.NewServiceContext(c)},
		},
	)

	ds.Start()
}
