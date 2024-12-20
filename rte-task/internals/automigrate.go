package internals

import (
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

func Automigration() {
	Connection := GetConnection()
	err := Connection.AutoMigrate(&models.UserDetails{}, &models.JobCreation{}, &models.UserJobDetails{})
	if err != nil {
		loggers.ErrorData.Println(err)
		return
	}
	loggers.InfoData.Println("Migrated tables Sucessfully")
}
