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

func MatchUserType(c *gin.Context, userId string) (err error) {
	userType := c.GetString("role_type")
	id := c.GetString("id")
	if userType == "USER" && id != userId {
		err = errors.New("unauthorized To access the resources")
		return err
	}

	err = CheckuserType(c, userType)
	return err
}
