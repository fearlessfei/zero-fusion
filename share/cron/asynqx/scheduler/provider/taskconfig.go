package provider

import (
	"zero-fusion/share/model"
)

// TaskConfig 任务配置
type TaskConfig struct {
	CronSpec    string             `json:"cron_spec"`    // cron表达式
	TaskType    string             `json:"task_type"`    // 任务类型
	Payload     []byte             `json:"payload"`      // 任务参数
	TaskOptions *model.TaskOptions `json:"task_options"` // 任务选项
}

type PeriodicTaskConfigContainer struct {
	Configs []*TaskConfig
}
