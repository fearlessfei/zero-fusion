package config

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"

	"zero-fusion/share/orm/gormx"
)

type Config struct {
	rest.RestConf

	OTELMeter  bool
	Gorm       gormx.Config
	CacheRedis redis.RedisConf
	BizRedis   redis.RedisConf
}
