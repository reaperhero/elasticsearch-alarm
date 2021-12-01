package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
	"github.com/sirupsen/logrus"
)

type DbRepo interface {
	CreateAlarmConfig(ac *model.AlarmConfig) error
	GetAlarmConfig(page, size int) (configs []model.AlarmConfig, err error)
}

type dbRepo struct {
	db *gorm.DB
}

func NewDbRepo() DbRepo {
	return &dbRepo{
		db: db,
	}
}

func (d *dbRepo) CreateAlarmConfig(ac *model.AlarmConfig) error {
	return d.db.Save(ac).Error
}

func (d *dbRepo) GetAlarmConfig(page, size int) (configs []model.AlarmConfig, err error) {
	if err = d.db.Find(&configs).Offset(page * size).Limit(size).Error; err == gorm.ErrRecordNotFound {
		logrus.Infof("[dbRepo.GetAlarmConfig] %s", err)
		return nil, nil
	}
	return
}
