package initialize

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/plugin/opentelemetry/tracing"

	"zero-fusion/share/orm/gormx"
)

func InitGormDB(config gormx.Config) *gormx.GormDB {
	gormDB := gormx.NewGormDB(config)
	if err := gormDB.Use(tracing.NewPlugin()); err != nil {
		logx.Must(fmt.Errorf("use trace plugin err: %v", err))
	}

	if config.Debug {
		gormDB.Debug()
	}

	return gormDB
}
