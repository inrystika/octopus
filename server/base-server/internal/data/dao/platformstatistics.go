package dao

import (
	"server/base-server/internal/data/dao/model"
	"server/common/constant"
	"server/common/errors"

	"gorm.io/gorm"
)

type PlatformStatisticsDao interface {
	Summary() (*model.PlatformStatSummary, error)
}

type platformStatisticsDao struct {
	db *gorm.DB
}

func NewPlatformStatisticsDao(db *gorm.DB) PlatformStatisticsDao {
	return &platformStatisticsDao{
		db: db,
	}
}

func (d platformStatisticsDao) Summary() (*model.PlatformStatSummary, error) {
	db := d.db
	var running int64
	res := db.Where("status = ? ", constant.RUNNING).Model(&model.TrainJob{}).Count(&running)
	if res.Error != nil {
		return nil, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	var pending int64
	res = db.Where("status = ? ", constant.PENDING).Model(&model.TrainJob{}).Count(&pending)
	if res.Error != nil {
		return nil, errors.Errorf(res.Error, errors.ErrorDBCountFailed)
	}

	return &model.PlatformStatSummary{TrainJob: &model.PlatformStatTrainJob{
		PendingNum: int(pending),
		RunningNum: int(running),
	}}, nil
}
