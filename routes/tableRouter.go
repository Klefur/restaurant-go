package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-restaurant/controllers"
)

func TableRoutes(router *gin.Engine) {

	router.GET("/tables", controller.GetTables())
	router.GET("/tables/:id", controller.GetTable())
	router.POST("/tables", controller.CreateTable())
	router.PATCH("/tables/:id", controller.UpdateTable())
}