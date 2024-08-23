package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-restaurant/controllers"
)

func UserRoutes(router *gin.RouterGroup) {

	router.GET("/users", controller.GetUsers())
	router.GET("/users/:id", controller.GetUser())
	router.POST("/users/signup", controller.SignUp())
	router.POST("/users/login", controller.Login())
}