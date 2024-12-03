package internals

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

var (
	GlobalConnection *gorm.DB
)

func Automigration() {
	err := GlobalConnection.AutoMigrate(&models.UserDetails{}, &models.JobCreation{}, &models.UserJobDetails{})
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
}
