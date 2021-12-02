package repository

import "github.com/reaperhero/elasticsearch-alarm/pkg/model"

type ElasticsearchRepo interface {
	SearchMessageWithText(request model.SearchRequestBody, limit int) ([][]byte, error)
}
