package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-restaurant/models"
)

func GetOrders() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orders []models.Order

		err := db.Find(&orders).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ "success": false, "error": "error occurred while fetching orders" })
			return
		}

		c.JSON(http.StatusOK, gin.H{ "success": true, "payload": orders })
	}
}

func GetOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		
		orderId := c.Param("id")
		var order models.Order

		err := db.Find(&order, orderId).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ "success": false, "error": "error occurred while fetching order" })
			return
		}

		if order.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{ "success": false, "error": "order was not found" })
			return
		}

		c.JSON(http.StatusOK, gin.H{ "success": true, "payload": order })
	}
}

func CreateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {

		var order models.Order

		err := c.BindJSON(&order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{ "success": false, "error": err.Error() })
			return
		}

		var table models.Table

		err = db.Find(&table, order.Table_id).Error
		if err != nil || table.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{ "success": false, "error": "table was not found" })
			return
		}

		err = db.Create(&order).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ "success": false, "error": "error occurred while creating order" })
			return
		}

		c.JSON(http.StatusOK, gin.H{ "success": true, "payload": order })
	}
}

func UpdateOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		var order models.Order

		err := c.BindJSON(&order)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{ "success": false, "error": err.Error() })
			return
		}

		orderId := c.Param("id")
		err = db.Model(&order).Where("id = ?", orderId).Updates(&order).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{ "success": false, "error": "error occurred while updating order" })
			return
		}

		c.JSON(http.StatusOK, gin.H{ "success": true, "payload": order })
	}
}
