package helpers

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

func CheckuserType(c *gin.Context, role string) error {
	usertype := c.GetString("role_type")
	if usertype != role {
		return fmt.Errorf("unauthorized to access the resource")
	}
	return nil
}

// check their Admin JobPosts with Fields
func CheckRoleType(post models.JobCreation, tokentype string, tokenid int, parmid int) error {
	if tokentype != "ADMIN" {
		return fmt.Errorf("invalid user-User have not access to view this details")
	}
	if tokenid != parmid {
		return fmt.Errorf("invalid ID,Your Payload ID and RoleId is Mismatching Here,Check It")
	}
	if post.DomainID != parmid {
		return fmt.Errorf("invalid ID,Your Payload ID and UserId is Mismatching Here,Check It")
	}
	return nil
}

// check their roles by admin or users
func ValidationCheckForRoleType(tokentype string, tokenid int, parmid int) error {
	if tokentype != "ADMIN" {
		return fmt.Errorf("invalid user-User have not access to view this details")
	}
	if tokenid != parmid {
		return fmt.Errorf("invalid ID,Your Payload ID and RoleId is Mismatching Here,Check It")
	}
	return nil
}
