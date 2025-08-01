package initialize

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
)

func InitBizRedis(conf redis.RedisConf) *redis.Redis {
	return redis.MustNewRedis(redis.RedisConf{
		Host: conf.Host,
		Pass: conf.Pass,
		Type: conf.Type,
	})
}

func InitCacheRedis(conf redis.RedisConf) *redis.Redis {
	return redis.MustNewRedis(redis.RedisConf{
		Host: conf.Host,
		Pass: conf.Pass,
		Type: conf.Type,
	})
}
