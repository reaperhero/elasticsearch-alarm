package repository

import (
	"context"
	"errors"
	elasv6 "github.com/olivere/elastic/v6"
	"github.com/reaperhero/elasticsearch-alarm/pkg/model"
	"log"
	"os"
	"time"
)

type elasticsearchRepo6 struct {
	client *elasv6.Client
}

func NewElasticsearchClientV6(instance model.AlarmInstance) ElasticsearchRepo {
	opts := []elasv6.ClientOptionFunc{
		elasv6.SetURL(instance.EsUrl),
		elasv6.SetBasicAuth(instance.EsUser, instance.EsPass),
	}
	opts = append(opts, elasv6.SetGzip(true),
		elasv6.SetHealthcheckInterval(20*time.Second),
		elasv6.SetErrorLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elasv6.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	client, _ := elasv6.NewClient(opts...)
	return &elasticsearchRepo6{client: client}
}

func (es elasticsearchRepo6) SearchMessageWithText(request model.SearchRequestBody, limit int) ([][]byte, error) {
	if request.Interval < time.Second*10 {
		request.Interval = time.Second * 60
	}
	start := time.Now().UTC().Add(request.Interval * -1).Format("2006-01-02T15:04:05.999Z")
	end := time.Now().UTC().Format("2006-01-02T15:04:05.999Z")
	rangeQuery := elasv6.NewRangeQuery("@timestamp").Gte(start).Lte(end).Format("strict_date_optional_time")
	boolQuery := elasv6.NewBoolQuery().Filter(rangeQuery)

	if request.FieldK != "" && request.FieldV != "" {
		termQuery := elasv6.NewTermQuery(request.FieldK, request.FieldV)
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
		messages = append(messages, *hit.Source)
	}
	return messages, nil
}
