package database

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getDBConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database")
		return nil
	}

	return db
}

var dbUser string = os.Getenv("DB_USER")
var dbPassword string = os.Getenv("DB_PASSWORD")
var dbName string = os.Getenv("DB_NAME")
var dbHost string = os.Getenv("DB_HOST")
var dsn string = "postgresql://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName
