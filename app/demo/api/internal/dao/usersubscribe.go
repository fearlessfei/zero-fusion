package dao

import (
	"context"
	"errors"

	"gorm.io/gorm"

	"zero-fusion/app/demo/api/internal/dao/query"
	"zero-fusion/app/demo/api/internal/model"
)

type UserSubscribeDao struct{}

func (u *UserSubscribeDao) GetByID(ctx context.Context, id int64) (*model.Subscribe, error) {
	subscribe, err := query.Subscribe.WithContext(ctx).GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	return subscribe, nil
}
