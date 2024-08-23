package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-restaurant/controllers"
)

func OrderRoutes(router *gin.RouterGroup) {
	
	router.GET("/orders", controller.GetOrders())
	router.GET("/orders/:id", controller.GetOrder())
	router.POST("/orders", controller.CreateOrder())
	router.PATCH("/orders/:id", controller.UpdateOrder())
}