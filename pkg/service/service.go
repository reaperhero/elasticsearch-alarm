package service

import (
	"github.com/reaperhero/elasticsearch-alarm/pkg/dto"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
)

type WebService interface {
	CreateAlarmConfig(config dto.DtoAlarmConfig) error
	GetAlarmConfig(page, size int) (configs []model.AlarmConfig, err error)
	DeleteAlarmConfig(id int) (err error)
	UpdateAlarmConfigById(id int, dtoConfig dto.DtoAlarmConfig) (err error)
	CreateAlarmInstance(instance dto.DtoAlarmInstance) error
	DeleteAlarmInstance(id int) error
	ListAlarmInstance(page, size int) (ins []model.AlarmInstance)
	UpdateAlarmInstance(id int, instance dto.DtoAlarmInstance) error
	MonitorElasticsearchAlarm() error
}
