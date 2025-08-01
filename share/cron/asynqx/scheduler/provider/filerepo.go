package provider

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/hibiken/asynq"
)

// FileBasedConfigProvider implements asynq.PeriodicTaskConfigProvider interface.
type FileBasedConfigProvider struct {
	filename string
}

func (p *FileBasedConfigProvider) GetAll() ([]*asynq.PeriodicTaskConfig, error) {
	data, err := os.ReadFile(p.filename)
	if err != nil {
		return nil, err
	}

	var c PeriodicTaskConfigContainer
	if err := yaml.Unmarshal(data, &c); err != nil {
		return nil, err
	}

	var configs []*asynq.PeriodicTaskConfig
	for _, cfg := range c.Configs {
		configs = append(configs, &asynq.PeriodicTaskConfig{Cronspec: cfg.CronSpec,
			Task: asynq.NewTask(cfg.TaskType, cfg.Payload)})
	}

	return configs, nil
}
