package svc

import (
	"zero-fusion/app/demo/api/internal/config"
	"zero-fusion/app/demo/api/internal/config/centerconfig"
	"zero-fusion/app/demo/api/internal/dao/query"
	"zero-fusion/share/configcenter/subscriber/naocs"
	"zero-fusion/share/initialize"
	"zero-fusion/share/orm/gormx"

	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/core/configcenter"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
)

type ServiceContext struct {
	Config config.Config

	// gorm
	GormDB *gormx.GormDB

	// redis
	BizRedis   *redis.Redis
	CacheRedis *redis.Redis

	// 配置中心
	ConfigCenter configurator.Configurator[centerconfig.CenterConfig]
}

func NewServiceContext(c config.Config) *ServiceContext {
	svc := &ServiceContext{
		Config: c,
	}

	initialize.InitOTELMeter(c.OTELMeter)

	svc.GormDB = initialize.InitGormDB(c.Gorm)
	query.SetDefault(svc.GormDB.DB)

	svc.BizRedis = initialize.InitBizRedis(c.BizRedis)
	svc.CacheRedis = initialize.InitCacheRedis(c.CacheRedis)

	svc.ConfigCenter = initialize.InitNacosConfig[centerconfig.CenterConfig](naocs.Config{
		SC: []constant.ServerConfig{
			*constant.NewServerConfig("127.0.0.1", 8848, constant.WithContextPath("/nacos")),
		},
		CC: constant.NewClientConfig(
			constant.WithNamespaceId(""),
			constant.WithTimeoutMs(5000),
			constant.WithNotLoadCacheAtStart(true),
			constant.WithLogDir("/tmp/nacos/log"),
			constant.WithCacheDir("/tmp/nacos/cache"),
			constant.WithLogLevel("debug"),
		),
		DataId: "test-data-2",
		Group:  "test-group",
	})

	getConfig, err := svc.ConfigCenter.GetConfig()
	if err != nil {
		panic(err)
	}
	logx.Infof("config: %s", getConfig)

	return svc
}
