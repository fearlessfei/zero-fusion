package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"

	"zero-fusion/app/demo/api/internal/dao/query"
	"zero-fusion/app/demo/api/scripts/internal/config"
	"zero-fusion/share/initialize"
	"zero-fusion/share/orm/gormx"
)

type ServiceContext struct {
	Config config.Config

	// gorm
	GormDB *gormx.GormDB

	// redis
	BizRedis   *redis.Redis
	CacheRedis *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config: c,
	}

	initialize.InitOTELMeter(c.OTELMeter)
	svc.GormDB = initialize.InitGormDB(c.Gorm)
	svc.BizRedis = initialize.InitBizRedis(c.BizRedis)
	svc.CacheRedis = initialize.InitCacheRedis(c.CacheRedis)

	query.SetDefault(svc.GormDB.DB)

	return svc
}
