package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"

	"zero-fusion/share/orm/gormx"
)

type Config struct {
	service.ServiceConf

	OTELMeter  bool
	Gorm       gormx.Config
	CacheRedis redis.RedisConf
	BizRedis   redis.RedisConf
}
