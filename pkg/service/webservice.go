package service

import "github.com/reaperhero/elasticsearch-alarm/pkg/repository"

type webService struct {
	dbRepo repository.DbRepo
	esRepo map[string]repository.ElasticsearchRepo
}

func NewWebService() WebService {
	return webService{
		dbRepo: repository.NewDbRepo(),
		esRepo: make(map[string]repository.ElasticsearchRepo, 0),
	}
}
