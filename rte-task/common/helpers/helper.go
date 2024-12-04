package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckuserType(c *gin.Context, role string) (err error) {
	usertype := c.GetString("role_type")
	if usertype != role {
		err = errors.New("unauthorized to access the resource")
		return err
	}
	return nil
}

// func BothUserType(c *gin.Context, role string) (err error) {
// 	usertype := c.GetString("role_type")
// 	// userType := c.GetString("role_type")
// 	if usertype != role {
// 		err = errors.New("unauthorized to access the resource")
// 		return err
// 	}
// 	return nil
// }
