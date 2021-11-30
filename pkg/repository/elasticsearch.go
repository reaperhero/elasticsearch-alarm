package repository

import (
	"github.com/olivere/elastic/v7"
	"log"
	"os"
	"time"
)

type ElasticsearchRepo interface {
}

type elasticsearchRepo struct {
	client *elastic.Client
}

func NewElasticsearchClient(opts ...elastic.ClientOptionFunc) ElasticsearchRepo {
	opts = append(opts, elastic.SetGzip(true),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetErrorLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	client, _ := elastic.NewClient(opts...)
	return &elasticsearchRepo{client: client}
}


