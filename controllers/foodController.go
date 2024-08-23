package controllers

import (
	"go-restaurant/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetFoods() gin.HandlerFunc {
	return func(c *gin.Context) {
		var foods []models.Food

		page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
		if err != nil || page < 1 {
			page = 1
		}
		limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
		if err != nil || limit < 1 {
			limit = 10
		}

		offset := (page - 1) * limit

		err = db.Offset(offset).Limit(limit).Find(&foods).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occurred while fetching foods"})
			return
		}

		count := new(int64)

		err = db.Model(&models.Food{}).Count(count).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occurred while fetching foods"})
			return
		}

		numPages := *count / int64(limit)

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": gin.H{"num_pages": numPages, "actual_page": page, "foods": foods}})
	}
}

func GetFood() gin.HandlerFunc {
	return func(c *gin.Context) {

		foodId := c.Param("id")
		var food models.Food

		err := db.Find(&food, foodId).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while fetching food"})
			return
		}

		if food.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "food was not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": food})
	}
}

func CreateFood() gin.HandlerFunc {
	return func(c *gin.Context) {

		var food models.Food

		err := c.BindJSON(&food)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		var menu models.Menu

		err = db.First(&menu, food.Menu_id).Error
		if err != nil || menu.ID == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": "menu was not found"})
			return
		}

		err = db.Create(&food).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while creating food"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": food})
	}
}

func UpdateFood() gin.HandlerFunc {
	return func(c *gin.Context) {
		var food models.Food
		var dbFood models.Food

		err := c.BindJSON(&food)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		foodId := c.Param("id")

		err = db.Find(&dbFood, foodId).Error
		if err != nil || food.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "food was not found"})
			return
		}

		if food.Name != "" && len(food.Name) > 3 {
			dbFood.Name = food.Name
		}

		if food.Price != 0 && food.Price > 0 {
			dbFood.Price = food.Price
		}

		if food.Food_image != "" {
			dbFood.Food_image = food.Food_image
		}

		err = db.Model(&dbFood).Where("id = ?", foodId).Updates(&dbFood).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while updating food"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": food})
	}
}
