package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func CheckuserType(c *gin.Context, role string) (err error) {
	usertype := c.GetString("role_type")
	// userType := c.GetString("role_type")
	if usertype != role {
		err = errors.New("unauthorized to access the resource")
		return err
	}
	return nil
}

func MatchUserType(c *gin.Context, userId int) (err error) {
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")
	if userType == "USER" && userid != userId {
		err = errors.New("unauthorized To access the resources")
		return err
	}

	err = CheckuserType(c, userType)
	return err
}
