package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go-restaurant/models"
)

func GetMenus() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menus []models.Menu

		err := db.Preload("Foods").Find(&menus).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while fetching menus"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": menus})
	}
}

func GetMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		menuId := c.Param("id")
		var menu models.Menu

		err := db.Preload("Foods").Find(&menu, menuId).Error

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occurred while fetching menu"})
			return
		}

		if menu.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "food was not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": menu})
	}
}

func CreateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu

		err := c.BindJSON(&menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		err = db.Create(&menu).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while creating menu"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": menu})
	}
}

func UpdateMenu() gin.HandlerFunc {
	return func(c *gin.Context) {
		var menu models.Menu
		var dbMenu models.Menu

		err := c.BindJSON(&menu)
		fmt.Println(menu)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		menuId := c.Param("id")
		err = db.Find(&dbMenu, menuId).Error
		fmt.Println(dbMenu)
		if err != nil || menu.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "menu was not found"})
			return
		}

		if menu.Name != "" && len(menu.Name) > 3 {
			dbMenu.Name = menu.Name
		}

		if menu.Category != "" && len(menu.Category) > 3 {
			dbMenu.Category = menu.Category
		}

		err = db.Model(&dbMenu).Where("id = ?", menuId).Updates(&dbMenu).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while updating menu"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": menu})
	}
}
