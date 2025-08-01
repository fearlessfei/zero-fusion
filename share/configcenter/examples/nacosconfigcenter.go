package main

import (
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/zeromicro/go-zero/core/logx"

	"zero-fusion/share/configcenter/subscriber/naocs"
	"zero-fusion/share/initialize"
)

type centerConfig struct{}

func main() {
	nacosConfig := initialize.InitNacosConfig[centerConfig](naocs.Config{
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
		DataId: "test-data",
		Group:  "test-group",
	})

	getConfig, err := nacosConfig.GetConfig()
	if err != nil {
		panic(err)
	}
	logx.Infof("config: %s", getConfig)
}
