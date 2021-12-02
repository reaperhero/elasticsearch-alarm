package cmd

import (
	"github.com/reaperhero/elasticsearch-alarm/handler/http"
	"github.com/reaperhero/elasticsearch-alarm/pkg/service"
	"github.com/sirupsen/logrus"
)

func Run() {
	service := service.NewWebService()
	go monitor(service)
	http.RunHttpserver(service)
}

func monitor(s service.WebService) {
	select {
	case msg := <-s.ElasticsConfigMonitor():
		logrus.Info(string(msg))
	}
}
