package helpers

import (
	"fmt"
)

// check their roles by users
func ValidateUserType(roleType string) error {
	if roleType != "USER" {
		return fmt.Errorf("invalid admin-Admin have not to access this details")
	}
	return nil
}

// check their roles by admin or users
func ValidateRoleType(roleType string) error {
	if roleType != "ADMIN" {
		return fmt.Errorf("invalid user-User have not access to view this details")
	}
	return nil
}
