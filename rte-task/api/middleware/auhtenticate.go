package middleware

import (
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/gin-gonic/gin"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		clientToken := c.Request.Header.Get("token")
		if clientToken == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid token generating"})
			c.Abort()
			return
		}

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
