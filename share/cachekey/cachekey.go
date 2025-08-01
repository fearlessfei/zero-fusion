package cachekey

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

const (
	defaultNameSpace = ""  // 命名空间
	defaultKeyPrefix = ""  // Key 前缀
	defaultSeparator = ":" // Key 分隔符
)

type Option func(*cacheKey)

// CacheKey 表示一个缓存 Key
type cacheKey struct {
	NameSpace string // key 命名空间
	KeyPrefix string // Key 前缀
	Separator string // Key 分隔符
	KeyParts  []any  // Key 部分
}

var cacheKeyPool = sync.Pool{
	New: func() any {
		return &cacheKey{}
	},
}

// putCacheKey CacheKey归还到池中
func putCacheKey(ck *cacheKey) {
	ck.NameSpace = ""
	ck.KeyPrefix = ""
	ck.Separator = ""
	ck.KeyParts = ck.KeyParts[:0]

	cacheKeyPool.Put(ck)
}

func newCacheKey(opts ...Option) *cacheKey {
	ck := cacheKeyPool.Get().(*cacheKey)
	ck.NameSpace = defaultNameSpace
	ck.KeyPrefix = defaultKeyPrefix
	ck.Separator = defaultSeparator
	ck.KeyParts = ck.KeyParts[:0]

	for _, opt := range opts {
		opt(ck)
	}

	return ck
}

// setKeyParts 设置 Key 部分
func (ck *cacheKey) setKeyParts(keyParts []any) {
	ck.KeyParts = append(ck.KeyParts[:0], keyParts...)
}

// buildKey 构造最终 Key，统一添加命名空间和前缀
func (ck *cacheKey) buildKey() string {
	if ck.KeyPrefix != "" {
		ck.KeyParts = append([]any{ck.KeyPrefix}, ck.KeyParts...)
	}

	if ck.NameSpace != "" {
		ck.KeyParts = append([]any{ck.NameSpace}, ck.KeyParts...)
	}

	var keys []string
	for _, key := range ck.KeyParts {
		keys = append(keys, fmt.Sprint(key))
	}

	return strings.Join(keys, ck.Separator)
}

// WithNameSpace 设置命名空间
func WithNameSpace(nameSpace string) Option {
	return func(ck *cacheKey) {
		ck.NameSpace = nameSpace
	}
}

// WithKeyPrefix 设置 Key 前缀
func WithKeyPrefix(keyPrefix string) Option {
	return func(ck *cacheKey) {
		ck.KeyPrefix = keyPrefix
	}
}

// WithSeparator 设置 Key 分隔符
func WithSeparator(separator string) Option {
	return func(ck *cacheKey) {
		ck.Separator = separator
	}
}

// BuildKey 构造最终 Key
func BuildKey(KeyParts []any, opts ...Option) string {
	ck := newCacheKey(opts...)
	defer putCacheKey(ck)

	ck.setKeyParts(KeyParts)
	key := ck.buildKey()
	return key
}

// KeyMeta 表示一个 Key 及其元信息
type KeyMeta struct {
	Key     string        // 实际 key
	TTL     time.Duration // 缓存时间，0 表示永久
	Desc    string        // 描述用途
	Version uint8         // Key 版本号（可选）
}

// Clone 克隆key
func (km *KeyMeta) Clone() *KeyMeta {
	return &KeyMeta{
		Key:     km.Key,
		TTL:     km.TTL,
		Desc:    km.Desc,
		Version: km.Version,
	}
}

// SetTTL 设置缓存时间
func (km *KeyMeta) SetTTL(duration time.Duration) {
	km.TTL = duration
}
