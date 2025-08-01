### 缓存key
- localkey 目录下为本地缓存key（进程缓存）
- rediskey 目录下为 redis 缓存key

### localkey, rediskey统一使用cachekey.BuildKey生成
``` go
func main() {
    keyMeta := demoKey(123)
    fmt.Println("key", keyMeta.Key)
    fmt.Println("ttl", keyMeta.TTL)
    fmt.Println("desc", keyMeta.Desc)
    fmt.Println("version", keyMeta.Version)
    
    // redis缓存
    redis.String.Set(keyMeta.Key, "value", keyMeta.TTL)

    // 本地缓存复用redis key,反之亦然
    keyMeta2 := keyMeta.Clone()
    keyMeta2.SetTTL(10 * time.Second)
    // 本地缓存
    cache.Set(keyMeta2.Key, "value", keyMeta2.TTL)
    
    // output:
    // key test:prefix:ab:cd:123:ef
    // ttl 0
    // desc demo key
    // version 1
}

func demoKey(id int) *cachekey.KeyMeta {
	key := cachekey.BuildKey(
	    []any{"ab", "cd", id, "ef"},
        WithNameSpace("test"),
        WithKeyPrefix("prefix"),
        WithSeparator(":"), // 默认为 ":",如果使用":"分隔符可以不设置
	)

	return &cachekey.KeyMeta{
		Key:     key,
		TTL:     0,
		Desc:    "demo key",
		Version: 1,
	}
}
```