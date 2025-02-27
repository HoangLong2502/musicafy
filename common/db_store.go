package common

import (
	"gorm.io/gorm"
)

type store struct {
	db *gorm.DB
}

func NewStore(sqldb *gorm.DB) *store {
	return &store{
		db: sqldb,
	}
}
