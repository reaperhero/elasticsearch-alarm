package service

import (
	"github.com/reaperhero/elasticsearch-alarm/pkg/dto"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
)

type WebService interface {
	CreateAlarmConfig(config dto.DtoAlarmConfig) error
	GetAlarmConfig(page, size int) (configs []model.AlarmConfig, err error)
}
