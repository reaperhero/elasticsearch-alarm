package repository

import (
	"context"
	"errors"
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
		elasv7.SetBasicAuth(instance.EsUser, instance.EsPass),
	}
	opts = append(opts, elasv7.SetGzip(true),
		elasv7.SetHealthcheckInterval(10*time.Second),
		elasv7.SetErrorLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elasv7.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
		elasv7.SetSniff(false),
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

func (es *elasticsearchRepo7) SearchMessageWithText(request model.SearchRequestBody, limit int) ([][]byte, error) {
	if request.Interval < time.Second*10 {
		request.Interval = time.Second * 60
	}
	start := time.Now().UTC().Add(request.Interval * -1).Format("2006-01-02T15:04:05.999Z")
	end := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	rangeQuery := elasv7.NewRangeQuery("@timestamp").Gte(start).Lte(end).Format("strict_date_optional_time")
	boolQuery := elasv7.NewBoolQuery().Filter(rangeQuery)

	if request.FieldK != "" && request.FieldV != "" {
		termQuery := elasv7.NewTermQuery(request.FieldK, request.FieldV)
		boolQuery.Filter(termQuery)
	}

	response, err := es.client.Search().
		Index(request.IndexName).
		Query(boolQuery).
		Sort("@timestamp", true).
		Size(limit).
		Do(context.Background())

	if err != nil {
		return nil, err
	}

	hits := response.Hits.Hits

	if len(hits) == 0 {
		return nil, errors.New("no record")
	}

	var messages [][]byte

	for _, hit := range hits {
		messages = append(messages, hit.Source)
	}
	return messages, nil
}
