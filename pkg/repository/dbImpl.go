package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
	"github.com/sirupsen/logrus"
)

type DbRepo interface {
	CreateAlarmConfig(ac *model.AlarmConfig) error
	GetAlarmConfig(page, size int) (configs []model.AlarmConfig, err error)
	DeleteAlarmConfig(id int) error
	GetAlarmConfigById(id int) (config *model.AlarmConfig)
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

func (d *dbRepo) GetAlarmConfigById(id int) (*model.AlarmConfig) {
	config := model.AlarmConfig{}
	if err := d.db.Find(&config,id).Error; err != nil {
		logrus.Infof("[dbRepo.GetAlarmConfigById] %s", err)
		return nil
	}
	return &config
}

func (d *dbRepo) DeleteAlarmConfig(id int) error {
	if err := d.db.Where("id=?", id).Delete(new(model.AlarmConfig)).Error; err == gorm.ErrRecordNotFound {
		logrus.Errorf("[dbRepo.DeleteAlarmConfig] %s", err)
		return nil
	} else {
		return err
	}
}
