package service

import (
	"github.com/reaperhero/elasticsearch-alarm/pkg/dto"
	pkgerr "github.com/reaperhero/elasticsearch-alarm/pkg/errors"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
	"github.com/reaperhero/elasticsearch-alarm/pkg/repository"
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
