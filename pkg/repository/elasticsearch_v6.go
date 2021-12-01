package repository

import (
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
		elasv6.SetHealthcheckInterval(10*time.Second),
		elasv6.SetErrorLog(log.New(os.Stdout, "ELASTIC ", log.LstdFlags)),
		elasv6.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)),
	)
	client, _ := elasv6.NewClient(opts...)
	return &elasticsearchRepo6{client: client}
}
