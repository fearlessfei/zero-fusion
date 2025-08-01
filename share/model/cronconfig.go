package model

import (
	"database/sql/driver"
	"fmt"
	"time"

	jsoniter "github.com/json-iterator/go"
)

type TaskOptions struct {
	MaxRetry  int           `json:"max_retry"`  // 最大重试次数
	Queue     string        `json:"queue"`      // 队列名
	TaskID    string        `json:"task_id"`    // 任务ID
	Timeout   time.Duration `json:"timeout"`    // 超时时间
	Deadline  time.Time     `json:"deadline"`   // 截止时间
	Unique    time.Duration `json:"unique"`     // 唯一性
	ProcessAt time.Time     `json:"process_at"` // 处理时间
	ProcessIn time.Duration `json:"process_in"` // 处理延时
	Retention time.Duration `json:"retention"`  // 保留时间
	Group     string        `json:"group"`      //
}

type CronConfig struct {
	CronSpec    string       `gorm:"column:cron_spec" json:"cron_spec"`       // cron表达式
	TaskType    string       `gorm:"column:task_type" json:"task_type"`       // 任务类型
	Payload     string       `gorm:"column:payload" json:"payload"`           // 任务参数
	TaskOptions *TaskOptions `gorm:"column:task_options" json:"task_options"` // 任务选项
	CreatedAt   time.Time    `gorm:"column:created_at" json:"created_at"`     // 创建时间
	UpdatedAt   time.Time    `gorm:"column:updated_at" json:"updated_at"`     // 更新时间
}

func (CronConfig) TableName() string {
	return "cron_config"
}

func (o *TaskOptions) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var data []byte
	switch v := value.(type) {
	case []byte:
		data = v
	case string:
		data = []byte(v)
	default:
		return fmt.Errorf("task options数据错误: unexpected type %T", value)
	}

	return jsoniter.Unmarshal(data, o)
}

func (o TaskOptions) Value() (driver.Value, error) {
	return jsoniter.Marshal(o)
}
