package middleware

import (
	"fmt"
	"net/http"

	"attendance-app/helper"

	"github.com/gin-gonic/gin"
)

// Authentication middleware that authenticates the root ahead of it.the routes need token as there header to access
func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("No Authorization header provided")})
			c.Abort()
			return
		}

		claims, err := helper.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("rollno", claims.RollNo)

		c.Set("uid", claims.UserID)

		c.Next()

	}
}
