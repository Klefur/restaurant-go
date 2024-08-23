package controllers

import (
	"go-restaurant/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetOrderItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItems []models.OrderItem

		err := db.Find(&orderItems).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while fetching orderItems"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": orderItems})
	}
}

func GetOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderItemId := c.Param("id")
		var orderItem models.OrderItem

		err := db.Find(&orderItem, orderItemId).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occurred while fetching orderItem"})
			return
		}

		if orderItem.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "orderItem was not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": orderItem})
	}
}

func GetOrderItemsByOrder() gin.HandlerFunc {
	return func(c *gin.Context) {
		orderId := c.Param("id")
		var orderItems []models.OrderItem

		err := db.Find(&orderItems, orderId).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occurred while fetching orderItems"})
			return
		}

		if len(orderItems) == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "orderItems were not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": orderItems})
	}
}

func CreateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItem models.OrderItem
		var food models.Food
		var order models.Order

		err := c.BindJSON(&orderItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		err = db.Find(&food, orderItem.Food_id).Error
		if err != nil || food.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "food was not found"})
			return
		}

		err = db.Find(&order, orderItem.Order_id).Error
		if err != nil || order.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "order was not found"})
			return
		}

		err = db.Create(&orderItem).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while creating orderItem"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": orderItem})
	}
}

func UpdateOrderItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		var orderItem models.OrderItem
		var food models.Food
		var order models.Order

		err := c.BindJSON(&orderItem)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}
		if orderItem.Food_id != 0 {
			err = db.Find(&food, orderItem.Food_id).Error
			if err != nil || food.ID == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "food was not found"})
				return
			}
		}
		if orderItem.Order_id != 0 {
			err = db.Find(&order, orderItem.Order_id).Error
			if err != nil || order.ID == 0 {
				c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "order was not found"})
				return
			}
		}

		orderItemId := c.Param("id")
		err = db.Model(&orderItem).Where("id = ?", orderItemId).Updates(orderItem).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while updating orderItem"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": orderItem})
	}
}
