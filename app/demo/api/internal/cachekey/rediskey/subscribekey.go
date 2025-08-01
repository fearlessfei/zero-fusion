package rediskey

import (
	"zero-fusion/share/cachekey"
)

var SubscribeRedisKey = subscribeRedisKey{}

// subscribeRedisKey 订阅 Key
type subscribeRedisKey struct{}

func (k subscribeRedisKey) buildKey(KeyParts []any) string {
	return cachekey.BuildKey(
		KeyParts,
		cachekey.WithKeyPrefix("subscribe"),
	)
}

func (k subscribeRedisKey) HasSubscribed(openid string, subID int64) *cachekey.KeyMeta {
	keyParts := []any{"has", "subscribed", openid, subID}

	return &cachekey.KeyMeta{
		Key:  k.buildKey(keyParts),
		TTL:  0,
		Desc: "是否已订阅",
	}
}
