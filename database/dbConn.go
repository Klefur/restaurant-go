package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "go-restaurant/models"
)

var dsn string

func InitDB() {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dsn = "postgresql://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database")
		panic(err)
	}

	// Uncomment the following line to drop all tables
	err = db.Migrator().DropTable(&model.User{}, &model.Food{}, &model.Menu{}, &model.Table{}, &model.Order{}, &model.OrderItem{}, &model.Invoice{})
	if err != nil {
		fmt.Println("error dropping tables")
		panic(err)
	}

	err = db.AutoMigrate(&model.User{}, &model.Food{}, &model.Menu{}, &model.Table{}, &model.Order{}, &model.OrderItem{}, &model.Invoice{})

	if err != nil {
		fmt.Println("error migrating tables")
		panic(err)
	}

	fmt.Println("Database initialized & migrated")
}

func GetDB() (*gorm.DB) {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	dsn = "postgresql://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Logger = db.Logger.LogMode(4)

	if err != nil {
		log.Fatal("error connecting to database")
		return nil
	}

	return db
}
