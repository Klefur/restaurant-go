package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-restaurant/controllers"
)

func FoodRoutes(router *gin.RouterGroup) {

	router.GET("/foods", controller.GetFoods())
	router.GET("/foods/:id", controller.GetFood())
	router.POST("/foods", controller.CreateFood())
	router.PATCH("/foods/:id", controller.UpdateFood())
}