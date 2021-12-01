package service

import (
	"encoding/json"
	"github.com/reaperhero/elasticsearch-alarm/pkg/repository"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
)

func (w *webService) MonitorElasticsearchAlarm() error {
	ins := w.dbRepo.ListAlarmInstance(0, 100)
	for _, in := range ins {
		esUrl := in.EsUrl
		if checkEsVersion(esUrl) == "7" {
			w.esRepoMap[esUrl] = repository.NewElasticsearchClientV7(in)
		}
		if checkEsVersion(in.EsUrl) == "6" {
			w.esRepoMap[esUrl] = repository.NewElasticsearchClientV7(in)
		}
	}
	return nil
}

func checkEsVersion(url string) string {
	resp, err := http.Get(url)

	if err != nil {
		logrus.WithField("error", err).Error("[checkEsVersion]")
		return ""
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	version := struct {
		Number string `json:"number"`
	}{}
	err = json.Unmarshal(body, version)
	if err != nil {
		logrus.WithField("error", err).Error("[checkEsVersion]")
		return ""
	}
	vs := strings.Split(version.Number, ".")
	return vs[0]
}
