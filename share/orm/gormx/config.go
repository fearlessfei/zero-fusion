package gormx

import (
	"time"
)

// ConnConfig 链接配置
type ConnConfig struct {
	Dialect          string        `json:",default=mysql"` // 驱动类型
	MasterDefaultDSN string        `json:""`               // 主库默认地址
	MasterDSNS       []string      `json:""`               // 主库地址
	SlaveDSNS        []string      `json:",optional"`      // 从库地址
	MaxIdleConns     int           `json:",default=100"`   // 链接最大空闲数
	MaxOpenConns     int           `json:",default=200"`   // 链接最大开启数
	ConnMaxLifetime  time.Duration `json:",default=2h"`    // 链接最大存活时间
}

// LoggerConfig 日志配置
type LoggerConfig struct {
	LoggerType                string        `json:""`               // 日志记录器类型
	SlowThreshold             time.Duration `json:",default=200ms"` // 慢查询阈值
	LogLevel                  string        `json:",default=info"`  // 日志级别
	IgnoreRecordNotFoundError bool          `json:",default=false"` // 忽略记录不存在错误
}

// Config Gorm配置
type Config struct {
	Conn   ConnConfig
	Logger LoggerConfig

	Debug bool `json:",default=false"` // debug模式
}
