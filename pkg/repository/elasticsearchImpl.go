package repository

import (
	"github.com/sirupsen/logrus"
)

func (es *elasticsearchRepo) ListIndexNames() (result []string) {
	var err error
	result, err = es.client.IndexNames()
	if err != nil {
		logrus.WithField("error", err).Error("[elasticsearchRepo.ListIndexNames] err")
	}
	return result
}
