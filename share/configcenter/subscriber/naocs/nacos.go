package naocs

import (
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/common/logger"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"github.com/zeromicro/go-zero/core/configcenter/subscriber"
	"github.com/zeromicro/go-zero/core/logx"
)

func init() {
	logger.SetLogger(newNacosLogger())
}

type Config struct {
	SC     []constant.ServerConfig
	CC     *constant.ClientConfig
	DataId string
	Group  string
}

// nacosSubscriber is a subscriber that subscribes to nacos.
type nacosSubscriber struct {
	client config_client.IConfigClient
	config Config
}

// MustNewNacosSubscriber returns a nacos Subscriber, exits on errors.
func MustNewNacosSubscriber(config Config) subscriber.Subscriber {
	s, err := NewNacosSubscriber(config)
	logx.Must(err)
	return s
}

// NewNacosSubscriber returns a nacos Subscriber.
func NewNacosSubscriber(config Config) (subscriber.Subscriber, error) {
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ServerConfigs: config.SC,
			ClientConfig:  config.CC,
		},
	)

	if err != nil {
		return nil, err
	}

	return &nacosSubscriber{
		client: client,
		config: config,
	}, nil
}

// AddListener adds a listener to the subscriber.
func (s *nacosSubscriber) AddListener(listener func()) error {
	return s.client.ListenConfig(vo.ConfigParam{
		DataId: s.config.DataId,
		Group:  s.config.Group,
		OnChange: func(namespace, group, dataId, data string) {
			logx.Infof("config changed namespace:%s,group:%s,dataId:%s,content:%s",
				namespace, group, dataId, data)
			listener()
		},
	})
}

// Value returns the value of the subscriber.
func (s *nacosSubscriber) Value() (string, error) {
	config, err := s.client.GetConfig(vo.ConfigParam{
		DataId: s.config.DataId,
		Group:  s.config.Group,
	})
	if err != nil {
		return "", err
	}

	return config, nil
}
