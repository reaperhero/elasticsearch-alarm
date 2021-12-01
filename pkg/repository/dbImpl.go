package repository

import "github.com/jinzhu/gorm"

type DbRepo interface {
}

type dbRepo struct {
	db *gorm.DB
}

func NewDbRepo() DbRepo {
	return dbRepo{
		db: db,
	}
}


