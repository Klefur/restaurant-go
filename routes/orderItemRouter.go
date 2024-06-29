package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-restaurant/controllers"
)

func OrderItemRoutes(router *gin.Engine) {

	router.GET("/orderItems", controller.GetOrderItems())
	router.GET("/orderItems/:id", controller.GetOrderItem())
	router.GET("/orderItems/order/:id", controller.GetOrderItemsByOrder())
	router.POST("/orderItems", controller.CreateOrderItem())
	router.PATCH("/orderItems/:id", controller.UpdateOrderItem())
}