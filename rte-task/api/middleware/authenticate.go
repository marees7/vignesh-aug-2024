package middleware

import (
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/gin-gonic/gin"
)

// Authenticate the users with their given token
func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Token is required"})
			c.Abort()
			return
		}

		if len(clientToken) > 7 && clientToken[:7] == "Bearer " {
			clientToken = clientToken[7:]
		} else {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Token must be prefixed with 'Bearer '"})
			c.Abort()
			return
		}

		//Validate their token with their details
		claims, err := validation.ValidateToken(clientToken)
		if err != "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		c.Set("email", claims.Email)
		c.Set("name", claims.Name)
		c.Set("role_type", claims.RoleType)
		c.Set("user_id", claims.UserID)
	}
}
