package main

import (
	"go-restaurant/middleware"
	"go-restaurant/routes"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	port := os.Getenv("PORT")

	if port == "" {
		port = "80"
	}

	router := gin.New()
	router.Use(gin.Logger())

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "Welcome to the restaurant API",
			})
		})

		routes.UserRoutes(api)

		api.Use(middleware.Authentication())

		routes.FoodRoutes(api)
		routes.MenuRoutes(api)
		routes.TableRoutes(api)
		routes.OrderRoutes(api)
		routes.OrderItemRoutes(api)
		routes.InvoiceRoutes(api)
	}

	router.Run(":" + port)

}
