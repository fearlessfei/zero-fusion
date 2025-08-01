package initialize

import (
	"zero-fusion/share/configcenter/subscriber/naocs"

	"github.com/zeromicro/go-zero/core/configcenter"
)

func InitNacosConfig[T any](config naocs.Config) configurator.Configurator[T] {
	return configurator.MustNewConfigCenter[T](configurator.Config{Type: "json"},
		naocs.MustNewNacosSubscriber(config))
}
