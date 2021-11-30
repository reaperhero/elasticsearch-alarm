package main

import (
	"github.com/olivere/elastic/v7"
	"github.com/reaperhero/elasticsearch-alarm/pkg/repository"
	"github.com/reaperhero/elasticsearch-alarm/pkg/utils"
	"github.com/sirupsen/logrus"
	"os"
)

var (
	indexName  = utils.GetEnvWithDeafult("INDEX", "")
	filedName  = utils.GetEnvWithDeafult("FIELDK", "")
	filedValue = utils.GetEnvWithDeafult("FIELDV", "")
	esUrl      = utils.GetEnvWithDeafult("ES_URL", "")
	esUser     = utils.GetEnvWithDeafult("ES_USER", "")
	esPass     = utils.GetEnvWithDeafult("ES_PASS", "")
	INTERVAL   = utils.GetEnvIntWithDefault("INTERVAL", 60)
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
}

func main() {
	repository.NewElasticsearchClient(
		elastic.SetURL(esUrl),
		elastic.SetBasicAuth(esUser, esPass),
	)
}
