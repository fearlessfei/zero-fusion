package provider

import (
	"github.com/hibiken/asynq"
)

// ProviderRepo 代表配置一个数据仓库
type ProviderRepo interface {
	GetAll() ([]*TaskConfig, error)
}

// ConfigProvider implements asynq.PeriodicTaskConfigProvider interface.
type ConfigProvider struct {
	Repo ProviderRepo
}

// GetConfigs Parses the yaml file and return a list of PeriodicTaskConfigs.
func (p *ConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	allConfigs, err := p.Repo.GetAll()
	if err != nil {
		panic(err)
	}

	var configs []*asynq.PeriodicTaskConfig
	for _, cfg := range allConfigs {
		opts := make([]asynq.Option, 0)
		if cfg.TaskOptions != nil {
			if cfg.TaskOptions.MaxRetry > 0 {
				opts = append(opts, asynq.MaxRetry(cfg.TaskOptions.MaxRetry))
			}
			if cfg.TaskOptions.Queue != "" {
				opts = append(opts, asynq.Queue(cfg.TaskOptions.Queue))
			}
			if cfg.TaskOptions.TaskID != "" {
				opts = append(opts, asynq.TaskID(cfg.TaskOptions.TaskID))
			}
			if cfg.TaskOptions.Timeout > 0 {
				opts = append(opts, asynq.Timeout(cfg.TaskOptions.Timeout))
			}
			if !cfg.TaskOptions.Deadline.IsZero() {
				opts = append(opts, asynq.Deadline(cfg.TaskOptions.Deadline))
			}
			if cfg.TaskOptions.Unique > 0 {
				opts = append(opts, asynq.Unique(cfg.TaskOptions.Unique))
			}
			if !cfg.TaskOptions.ProcessAt.IsZero() {
				opts = append(opts, asynq.ProcessAt(cfg.TaskOptions.ProcessAt))
			}
			if cfg.TaskOptions.ProcessIn > 0 {
				opts = append(opts, asynq.ProcessIn(cfg.TaskOptions.ProcessIn))
			}
			if cfg.TaskOptions.Retention > 0 {
				opts = append(opts, asynq.Retention(cfg.TaskOptions.Retention))
			}
			if cfg.TaskOptions.Group != "" {
				opts = append(opts, asynq.Group(cfg.TaskOptions.Group))
			}
		}

		configs = append(configs, &asynq.PeriodicTaskConfig{Cronspec: cfg.CronSpec,
			Task: asynq.NewTask(cfg.TaskType, cfg.Payload, opts...)})
	}

	return configs, nil
}
