package cron

import (
	"context"

	"github.com/hibiken/asynq"
)

// TaskContext 任务上下文，用于传递任务上下文
type TaskContext interface {
	Value(key string) any
}

// Task 代表一个通用任务的接口
type Task interface {
	// Name 返回任务名称
	Name() string
	// Run 执行任务
	Run(ctx context.Context, payload []byte, taskContext TaskContext) error
}

// Scheduler 调度器接口，用于注册并启动任务
type Scheduler interface {
	// Register 注册任务
	Register(task Task)
	// Start 启动调度器
	Start() error
	// Stop 停止调度器
	Stop() error
}

// PeriodicTask 代表一个周期任务的接口
type PeriodicTask interface {
	// Name 返回任务名称
	Name() string
	// CronSpec 返回cron表达式
	CronSpec() string
	// Task 返回任务
	Task() *asynq.Task
	// Options 返回任务选项
	Options() []asynq.Option
}

// PeriodicScheduler 周期调度器接口
type PeriodicScheduler interface {
	// Register 注册任务
	Register(task PeriodicTask)
	// Start 启动调度器
	Start()
	// Stop 停止调度器
	Stop()
}

// PeriodicDynamicScheduler 动态周期调度器接口
type PeriodicDynamicScheduler interface {
	// Register 注册任务
	Register(provider asynq.PeriodicTaskConfigProvider)
	// Start 启动调度器
	Start()
	// Stop 停止调度器
	Stop()
}
