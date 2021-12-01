package service

import (
	"github.com/reaperhero/elasticsearch-alarm/pkg/dto"
	pkgerr "github.com/reaperhero/elasticsearch-alarm/pkg/errors"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
	"github.com/reaperhero/elasticsearch-alarm/pkg/repository"
	"github.com/reaperhero/elasticsearch-alarm/pkg/utils"
	"github.com/sirupsen/logrus"
)

type webService struct {
	dbRepo repository.DbRepo
	esRepo map[string]repository.ElasticsearchRepo
}

func NewWebService() WebService {
	return &webService{
		dbRepo: repository.NewDbRepo(),
		esRepo: make(map[string]repository.ElasticsearchRepo, 0),
	}
}

func (w *webService) CreateAlarmConfig(config dto.DtoAlarmConfig) error {
	if err := w.dbRepo.CreateAlarmConfig(&model.AlarmConfig{
		EsIndex:       config.EsIndex,
		MsgType:       config.MsgType,
		MsgDefine:     config.MsgDefine,
		CheckInterval: config.CheckInterval,
		IsRunning:     true,
		MailUser:      config.MailUser,
		DingToken:     config.DingToken,
		DingMobiles:   config.DingMobiles,
	}); err != nil {
		logrus.WithField("error", err).Errorf("[webService.CreateAlarmConfig] %s", err)
		return pkgerr.ErrDbOperation
	}
	return nil
}

func (w *webService) GetAlarmConfig(page, size int) (configs []model.AlarmConfig, err error) {
	configs, err = w.dbRepo.GetAlarmConfig(page, size)
	if err != nil {
		return nil, pkgerr.ErrDbOperation
	}
	return configs, nil
}

func (w *webService) DeleteAlarmConfig(id int) (err error) {
	err = w.dbRepo.DeleteAlarmConfig(id)
	if err != nil {
		return pkgerr.ErrDbOperation
	}
	return nil
}

func (w *webService) UpdateAlarmConfigById(id int, dtoConfig dto.DtoAlarmConfig) (err error) {
	config := w.dbRepo.GetAlarmConfigById(id)
	if config == nil {
		err = pkgerr.ErrDbRecord
		return
	}
	if err := utils.CopyFields(config, dtoConfig); err != nil {
		logrus.WithField("CopyFields", err).Error()
	}
	err = w.dbRepo.CreateAlarmConfig(config)
	if err != nil {
		err = pkgerr.ErrDbOperation
	}
	return
}

func (w *webService) CreateAlarmInstance(instance dto.DtoAlarmInstance) error {
	if i := w.dbRepo.GetInstanceByUrl(instance.EsUrl); i != nil {
		return pkgerr.ErrDbExistRecord
	}
	if err := w.dbRepo.CreateAlarmInstance(&model.AlarmInstance{
		EsName: instance.EsName,
		EsUrl:  instance.EsUrl,
		EsUser: instance.EsUser,
		EsPass: instance.EsPass,
	}); err != nil {
		logrus.WithField("error", err).Errorf("[webService.CreateAlarmInstance] %s", err)
		return pkgerr.ErrDbOperation
	}
	return nil
}

func (w *webService) DeleteAlarmInstance(id int) error {
	if w.dbRepo.GetAlarmConfigById(id) != nil {
		return pkgerr.ErrDbRecord
	}
	if err := w.dbRepo.DeleteAlarmInstance(id); err != nil {
		return pkgerr.ErrDbOperation
	}
	return nil
}

func (w *webService) ListAlarmInstance(page, size int) []model.AlarmInstance {
	ins := w.dbRepo.ListAlarmInstance(page, size)
	return ins
}

func (w *webService) UpdateAlarmInstance(id int, instance dto.DtoAlarmInstance) error {
	oldInstance := w.dbRepo.GetInstanceById(id)
	if oldInstance != nil {
		return pkgerr.ErrDbRecord
	}
	if err := utils.CopyFields(oldInstance, &instance); err != nil {
		logrus.WithField("CopyFields", err).Error()
		return pkgerr.ErrDbRecord
	}
	if err := w.dbRepo.SaveAlarmInstance(oldInstance); err != nil {
		return pkgerr.ErrDbOperation
	}
	return nil
}
