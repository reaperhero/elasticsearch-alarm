package repository

import (
	"context"
	elasv7 "github.com/olivere/elastic/v7"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
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