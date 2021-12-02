package service

import (
	"bytes"
	"encoding/json"
	"errors"
	pkgerr "github.com/reaperhero/elasticsearch-alarm/pkg/errors"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
	"github.com/reaperhero/elasticsearch-alarm/pkg/repository"
	"github.com/reaperhero/elasticsearch-alarm/pkg/utils"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

func (w *webService) ElasticsConfigMonitor() (queryChan chan []byte) {
	go func() {
		if err := w.initElasticsearchConnect(); err != nil {
			logrus.Errorf("[webService.ElasticsConfigMonitor]%s", err)
			return
		}
	}()
	conMap := w.dbRepo.ListAlarmInstanceWithConfig()
	for instance, configs := range conMap {
		conn := w.esRepoMap[instance.EsUrl]
		for _, config := range configs {
			go func(config model.AlarmConfig) {
				for {
					define := strings.Split(config.MsgDefine, ":")
					if len(define) != 2 {
						define[0] = ""
						define[1] = ""
					}
					requestBody := model.SearchRequestBody{
						IndexName: config.EsIndiceName,
						Interval:  time.Duration(config.CheckInterval) * time.Second,
						QType:     config.MsgType,
						FieldK:    define[0],
						FieldV:    define[1],
					}
					if response, err := conn.SearchMessageWithText(requestBody, 100); err != nil {
						logrus.WithField("error", err).Errorf("[conn.SearchMessageWithText] %s", instance.EsUrl)
					} else {
						for _, byteMsg := range response {
							queryChan <- byteMsg
						}
					}
					time.Sleep(time.Duration(config.CheckInterval) * time.Second)
				}
			}(config)
		}
	}
	return nil
}

func (w *webService) initElasticsearchConnect() error {
	conMap := w.dbRepo.ListAlarmInstanceWithConfig()
	if conMap == nil {
		return pkgerr.ErrDbRecord
	}
	for instance, _ := range conMap {
		esUrl := instance.EsUrl
		if checkEsVersion(esUrl) == "7" {
			w.esRepoMap[esUrl] = repository.NewElasticsearchClientV7(instance)
		}
		if checkEsVersion(esUrl) == "6" {
			w.esRepoMap[esUrl] = repository.NewElasticsearchClientV6(instance)
		}
	}
	for _, repo := range w.esRepoMap {
		if repo == nil {
			return errors.New("conn elastics err")
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

var (
	timeOut = utils.GetEnvInt64WithDefault("TIMEOUT", 3000) * 1000000
)

func handleSource(source []byte) (listMessage string) {
	if bytes.ContainsAny(source, "latency_human") {
		p := model.HistoryAccessPoint{}

		if err := json.Unmarshal(source, &p); err != nil {
			logrus.Errorf("[access Unmarshal err %s]", err)
			return
		}
		if err := p.CreateAmessage(); err != nil {
			return
		}
		if p.Amessage.Latency > timeOut {
			listMessage = "[request uri:" + p.Amessage.URI + ";latency_human:" + p.Amessage.LatencyHuman + "]"
		}
	}
	if bytes.ContainsAny(source, "level") {
		p := model.HistoryGrpcPoint{}
		if err := json.Unmarshal(source, &p); err != nil {
			logrus.Errorf("[grpc Unmarshal err %s]", err)
			return
		}
		if err := p.CreateGmessage(); err != nil {
			return
		}
		if p.Gmessage.Level == "error" {
			listMessage = "[request grpc:" + p.Gmessage.Msg + "]"
		}
	}
	if bytes.ContainsAny(source, "/usr/local/go/src/runtime/panic.go") {
		listMessage = "服务重启了"
	}

	return
}
