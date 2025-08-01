package asynqx

import (
	"github.com/hibiken/asynq"
)

// NewAsynqClient asynq客户端
func NewAsynqClient(clientOpt asynq.RedisClientOpt) *asynq.Client {
	return asynq.NewClient(clientOpt)
}
