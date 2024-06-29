package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"

	helper "go-restaurant/helpers"
)

// AuthMiddleware function
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{ "success": false, "error": "No token provided" })
			c.Abort()
			return
		}

		isValid, msg := helper.ValidateToken(token)
		if !isValid || msg != ""{
			c.JSON(http.StatusUnauthorized, gin.H{ "success": false, "error": msg })
			c.Abort()
			return
		}
		c.Next()
	}
}