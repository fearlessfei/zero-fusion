package main

import (
	"flag"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/conf"

	"zero-fusion/app/demo/api/crontask/worker/internal/config"
	"zero-fusion/app/demo/api/crontask/worker/internal/svc"
	"zero-fusion/app/demo/api/crontask/worker/register"
	"zero-fusion/share/cron/asynqx/worker"
	"zero-fusion/share/log"
)

var configFile = flag.String("f", "etc/dev.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	c.MustSetUp()

	log.OutStdout(c.Mode)

	ws := worker.NewWorkerServer(
		asynq.RedisClientOpt{
			Addr:     c.AsynqRedis.Host,
			Password: c.AsynqRedis.Pass,
		},
		asynq.Config{},
	)

	register.RegisterTask(ws, svc.NewServiceContext(c))

	ws.Start()
}
