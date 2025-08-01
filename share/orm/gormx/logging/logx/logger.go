package logx

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm/logger"
)

// Config logger配置
type Config struct {
	SlowThreshold             time.Duration
	LogLevel                  logger.LogLevel
	IgnoreRecordNotFoundError bool
}

// GormLogger 适配go-zero logx的grom日志封装
type GormLogger struct {
	Config
}

// NewGormLogger 创建一个适配go-zero的gorm日志器
func NewGormLogger(config Config) *GormLogger {
	return &GormLogger{
		Config: config,
	}
}

// LogMode 设置日志级别
func (l *GormLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.LogLevel = level
	return &newLogger
}

// Info 级别日志
func (l *GormLogger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Info {
		logx.WithContext(ctx).Infof("[gorm] "+msg, data...)
	}
}

// Warn 级别日志
func (l *GormLogger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Warn {
		logx.WithContext(ctx).Infof("[gorm] "+msg, data...)
	}
}

// Error 级别日志
func (l *GormLogger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.LogLevel >= logger.Error {
		logx.WithContext(ctx).Errorf("[gorm] "+msg, data...)
	}
}

// Trace 记录sql语句执行时间、慢查询及错误信息
func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	if l.LogLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin) // 计算执行时间
	sql, rows := fc()            // 获取sql语句和返回的行数

	// 记录sql语句
	logMsg := fmt.Sprintf("[sql] %s | rows: %d | elapsed: %s", sql, rows, elapsed)

	switch {
	case err != nil && l.LogLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound) || !l.IgnoreRecordNotFoundError):
		logx.WithContext(ctx).Errorf("%s | error: %v", logMsg, err)
	case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= logger.Warn:
		logx.WithContext(ctx).Slowf("%s | slow query (threshold: %s)", logMsg, l.SlowThreshold)
	default:
		logx.WithContext(ctx).Infof(logMsg)
	}
}
