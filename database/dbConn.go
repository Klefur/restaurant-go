package database

import (
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	db := getDBConnection()
	db.Logger = db.Logger.LogMode(4)

	return db
}
