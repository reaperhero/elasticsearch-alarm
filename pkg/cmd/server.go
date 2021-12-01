package cmd

import (
	"github.com/reaperhero/elasticsearch-alarm/handler/http"
	"github.com/reaperhero/elasticsearch-alarm/pkg/service"
)

func Run() {
	service := service.NewWebService()
	service.MonitorElasticsearchAlarm()
	http.RunHttpserver(service)
}
