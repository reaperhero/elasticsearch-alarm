package repository

import (
	"github.com/olivere/elastic/v7"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

type elasticsearchRepo7 struct {
	client *elastic.Client
}

func NewElasticsearchClientV7(opts ...elastic.ClientOptionFunc) ElasticsearchRepo {
	opts = append(opts, elastic.SetGzip(true),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	client, _ := elastic.NewClient(opts...)
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
