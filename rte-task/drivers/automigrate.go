package drivers

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/models"
	"gorm.io/gorm"
)

var (
	GlobalConnection *gorm.DB
)

func Automigration() {
	err := GlobalConnection.AutoMigrate(&models.UsersTable{}, &models.JobCreation{}, &models.UserJobDetails{})
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
}
