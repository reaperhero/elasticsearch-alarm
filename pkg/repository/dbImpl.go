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
	GetInstanceByUrl(url string) *model.AlarmInstance
	CreateAlarmInstance(ai *model.AlarmInstance) error
	DeleteAlarmInstance(id int) error
	GetInstanceById(id int) *model.AlarmInstance
	ListAlarmInstance(page, size int) (ins []model.AlarmInstance)
	SaveAlarmInstance(instance *model.AlarmInstance) error
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

func (d *dbRepo) GetAlarmConfigById(id int) *model.AlarmConfig {
	config := model.AlarmConfig{}
	if err := d.db.Find(&config, id).Error; err != nil {
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

func (d *dbRepo) GetInstanceByUrl(url string) *model.AlarmInstance {
	instance := model.AlarmInstance{}
	if err := d.db.Where("es_url=?", url).Find(&instance).Error; err != nil {
		logrus.Infof("[dbRepo.GetAlarmConfigById] %s", err)
		return nil
	}
	return &instance
}

func (d *dbRepo) GetInstanceById(id int) *model.AlarmInstance {
	instance := model.AlarmInstance{}
	if err := d.db.Find(&instance, id).Error; err != nil {
		return nil
	}
	return &instance
}

func (d *dbRepo) CreateAlarmInstance(ai *model.AlarmInstance) error {
	return d.db.Save(ai).Error
}

func (d *dbRepo) DeleteAlarmInstance(id int) error {
	err := d.db.Delete(new(model.AlarmInstance), id).Error
	if err != nil {
		logrus.Infof("[dbRepo.DeleteAlarmInstance] %s", err)
		return err
	}
	return nil
}

func (d *dbRepo) ListAlarmInstance(page, size int) (ins []model.AlarmInstance) {
	if err := d.db.Find(&ins).Offset(page * size).Limit(size).Error; err != nil {
		logrus.Infof("[dbRepo.ListAlarmInstance] %s", err)
	}
	return
}

func (d *dbRepo) SaveAlarmInstance(instance *model.AlarmInstance) error {
	if err := d.db.Save(instance).Error; err != nil {
		logrus.Infof("[dbRepo.SaveAlarmInstance] %s", err)
		return err
	}
	return nil
}
