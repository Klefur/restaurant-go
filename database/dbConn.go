package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "go-restaurant/models"
)

func InitDB() {

	godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "postgresql://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName,
	}), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect to database")
		panic(err)
	}

	// Uncomment the following line to drop all tables
	// err = conn.Migrator().DropTable(&model.User{}, &model.Food{}, &model.Menu{}, &model.Table{}, &model.Order{}, &model.OrderItem{}, &model.Invoice{})

	err = conn.AutoMigrate(&model.User{}, &model.Food{}, &model.Menu{}, &model.Table{}, &model.Order{}, &model.OrderItem{}, &model.Invoice{})

	if err != nil {
		fmt.Println("error migrating tables")
		panic(err)
	}

	fmt.Println("Database initialized & migrated")
}

func GetDB() *gorm.DB {

	godotenv.Load()

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: "postgresql://" + dbUser + ":" + dbPassword + "@" + dbHost + "/" + dbName,
	}), &gorm.Config{})

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
		os.Exit(1)
	}

	return db
}
