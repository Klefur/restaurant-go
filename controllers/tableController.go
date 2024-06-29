package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-restaurant/models"
)

func GetTables() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tables []models.Table

		err := db.Find(&tables).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ "success": false, "error": "error occurred while fetching tables" })
			return
		}

		c.JSON(http.StatusOK, gin.H{ "success": true, "payload": tables })
	}
}

func GetTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		tableId := c.Param("id")
		var table models.Table

		err := db.Find(&table, tableId).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ "success": false, "error": "error occurred while fetching table" })
			return
		}

		if table.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{ "success": false, "error": "table was not found" })
			return
		}

		c.JSON(http.StatusOK, gin.H{ "success": true, "payload": table })
	}
}

func CreateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table

		err := c.BindJSON(&table)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{ "success": false, "error": err.Error() })
			return
		}

		err = db.Create(&table).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ "success": false, "error": "error occurred while creating table" })
			return
		}

		c.JSON(http.StatusOK, gin.H{ "success": true, "payload": table })
	}
}

func UpdateTable() gin.HandlerFunc {
	return func(c *gin.Context) {
		var table models.Table

		err := c.BindJSON(&table)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{ "success": false, "error": err.Error() })
			return
		}

		err = db.Model(&table).Where("id = ?", table.ID).Updates(&table).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ "success": false, "error": "error occurred while updating table" })
			return
		}

		c.JSON(http.StatusOK, gin.H{ "success": true, "payload": table })
	}
}