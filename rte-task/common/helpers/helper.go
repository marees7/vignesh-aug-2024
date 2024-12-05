package helpers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func CheckuserType(c *gin.Context, role string) error {
	usertype := c.GetString("role_type")
	if usertype != role {
		return fmt.Errorf("unauthorized to access the resource")
	}
	return nil
}
