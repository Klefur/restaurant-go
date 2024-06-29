package main

import (
	"go-restaurant/database"
	"go-restaurant/middleware"
	"go-restaurant/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}
	
	database.InitDB()

	router := gin.New()
	router.Use(gin.Logger())
	
	routes.UserRoutes(router)

	router.Use(middleware.Authentication()) 

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)

}