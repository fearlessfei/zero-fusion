package config

import (
	"zero-fusion/share/orm/gormx"

	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type Config struct {
	service.ServiceConf

	OTELMeter  bool
	Gorm       gormx.Config
	AsynqRedis redis.RedisConf
	CacheRedis redis.RedisConf
	BizRedis   redis.RedisConf
}
