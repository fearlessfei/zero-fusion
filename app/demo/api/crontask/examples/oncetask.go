package main

import (
	"zero-fusion/app/demo/api/crontask"
	"zero-fusion/share/cron/asynqx"

	"github.com/hibiken/asynq"
	"github.com/zeromicro/go-zero/core/logx"
)

func main() {
	client := asynqx.NewAsynqClient(asynq.RedisClientOpt{
		Addr:     "127.0.0.1:6379",
		Password: "",
	})
	task := asynq.NewTask(crontask.TaskTypeOnceDemo, []byte("once payload"))

	taskInfo, err := client.Enqueue(task)
	if err != nil {
		panic(err)
	}

	logx.Infof("enqueue task %s success, taskID: %s", task.Type(), taskInfo.ID)
}
