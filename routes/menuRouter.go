package routes

import (
	"github.com/gin-gonic/gin"
	controller "go-restaurant/controllers"
)

func MenuRoutes(router *gin.RouterGroup) {
		
	router.GET("/menus", controller.GetMenus())
	router.GET("/menus/:id", controller.GetMenu())
	router.POST("/menus", controller.CreateMenu())
	router.PATCH("/menus/:id", controller.UpdateMenu())
}
