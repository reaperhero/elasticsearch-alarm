package repository

import (
	"bytes"
	"context"
	"encoding/json"
	elasv7 "github.com/olivere/elastic/v7"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
	"github.com/reaperhero/elasticsearch-alarm/pkg/utils"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

type elasticsearchRepo7 struct {
	client *elasv7.Client
}

func NewElasticsearchClientV7(instance model.AlarmInstance) ElasticsearchRepo {
	opts := []elasv7.ClientOptionFunc{
		elasv7.SetURL(instance.EsUrl),
		elasv7.SetBasicAuth(instance.EsUser,instance.EsPass),
	}
	opts = append(opts, elasv7.SetGzip(true),
		elasv7.SetHealthcheckInterval(10*time.Second),
		elasv7.SetErrorLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elasv7.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	client, _ := elasv7.NewClient(opts...)
	return &elasticsearchRepo7{client: client}
}

func (es *elasticsearchRepo7) ListIndexNames() (result []string) {
	var err error
	result, err = es.client.IndexNames()
	if err != nil {
		logrus.WithField("error", err).Error("[elasticsearchRepo.ListIndexNames] err")
	}
	return result
}


func (es *elasticsearchRepo7) SearchMessageWithField(rangeTime time.Duration, indexname string, queryfileds map[string]string) (count int64, messages []string) {
	indexname = indexname + "-" + time.Now().Format("2006.01.02")
	start := time.Now().UTC().Add(rangeTime * -1).Format("2006-01-02T15:04:05.999Z")
	end := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	rangeQuery := elasv7.NewRangeQuery("@timestamp").Gte(start).Lte(end).Format("strict_date_optional_time")
	boolQuery := elasv7.NewBoolQuery().Filter(rangeQuery)

	for k, v := range queryfileds {
		matchPhrase := elasv7.NewMatchPhraseQuery(k, v)
		boolQuery.Filter(matchPhrase)
	}

	searchResult, err := es.client.Search().
		Index(indexname).
		Query(boolQuery).
		From(0).
		Size(2000).
		Pretty(true).
		Do(context.Background())
	if err != nil {
		logrus.Info(err)
		return
	}
	searchList := []string{}
	for _, hit := range searchResult.Hits.Hits {
		hitString := e.handleSource(hit.Source)
		if hitString != "" {
			searchList = append(searchList, hitString)
		}
	}
	return searchResult.TotalHits(), searchList
}


var (
	// 3s = 3000000000
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
