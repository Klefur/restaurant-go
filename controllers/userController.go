package controllers

import (
	helper "go-restaurant/helpers"
	"go-restaurant/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var users []models.User

		err := db.Find(&users).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while fetching users"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": users})
	}
}

func GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("id")
		var user models.User

		err := db.Find(&user, userId).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while fetching user"})
			return
		}

		if user.ID == 0 {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "user was not found"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": user})
	}
}

func SignUp() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		user.Password = helper.HashPassword(user.Password)

		err = db.Create(&user).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": "error occured while creating user"})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": user})
	}
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user models.User

		err := c.BindJSON(&user)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
			return
		}

		var dbUser models.User

		err = db.Where("email = ?", user.Email).First(&dbUser).Error
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"success": false, "error": "user not found"})
			return
		}

		passwordIsValid, msg := helper.VerifyPassword(dbUser.Password, user.Password)
		if !passwordIsValid {
			c.JSON(http.StatusUnauthorized, gin.H{"success": false, "error": msg})
			return
		}

		token, err := helper.GenerateToken(dbUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "payload": token})
	}
}
