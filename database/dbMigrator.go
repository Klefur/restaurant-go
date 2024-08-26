package database

import (
	"fmt"

	model "go-restaurant/models"
)

func InitDB() {
	db := getDBConnection()

	err := db.Migrator().CreateTable(&model.User{}, &model.Food{}, &model.Menu{}, &model.Table{}, &model.Order{}, &model.OrderItem{}, &model.Invoice{})
	if err != nil {
		fmt.Println("error creating tables")
	}

	fmt.Println("Database initialized")
}

func DropTables() {
	db := getDBConnection()

	err := db.Migrator().DropTable(&model.User{}, &model.Food{}, &model.Menu{}, &model.Table{}, &model.Order{}, &model.OrderItem{}, &model.Invoice{})
	if err != nil {
		fmt.Println("error dropping tables")
	}
}

func MigrateDB() {
	db := getDBConnection()

	err := db.AutoMigrate(&model.User{}, &model.Food{}, &model.Menu{}, &model.Table{}, &model.Order{}, &model.OrderItem{}, &model.Invoice{})

	if err != nil {
		fmt.Println("error migrating tables")
		panic(err)
	}

	fmt.Println("Database migrated")
}
