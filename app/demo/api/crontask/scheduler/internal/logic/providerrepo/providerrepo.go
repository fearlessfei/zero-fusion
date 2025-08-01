package providerrepo

import (
	"zero-fusion/app/demo/api/crontask/scheduler/internal/svc"
	"zero-fusion/share/cron/asynqx/scheduler/provider"
	"zero-fusion/share/model"
)

type DBRepo struct {
	SvcCtx *svc.ServiceContext
}

func (r *DBRepo) GetAll() ([]*provider.TaskConfig, error) {
	var configs []*model.CronConfig

	tx := r.SvcCtx.GormDB.DB.Find(&configs)
	if tx.Error != nil {
		return nil, tx.Error
	}

	taskConfigs := make([]*provider.TaskConfig, 0, len(configs))
	for _, config := range configs {
		taskConfigs = append(taskConfigs, &provider.TaskConfig{
			CronSpec:    config.CronSpec,
			TaskType:    config.TaskType,
			Payload:     []byte(config.Payload),
			TaskOptions: config.TaskOptions,
		})
	}

	return taskConfigs, nil
}
