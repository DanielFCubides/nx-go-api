package sql

import "github.com/jinzhu/gorm"

type Connection interface {
	GetDatabase() *gorm.DB
}
