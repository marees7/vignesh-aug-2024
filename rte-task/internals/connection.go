package internals

import (
	"fmt"
	"os"

	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ConnectionNew struct {
	*gorm.DB
}

func ConnectingDatabase() *ConnectionNew {
	path := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PORT"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DBNAME"))
	Connection, err := gorm.Open(postgres.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	defer HandlePanic()
	loggers.InfoData.Println("Connected sucessfully")
	return &ConnectionNew{
		Connection,
	}
}

func HandlePanic() {
	if err := recover(); err != nil {
		loggers.ErrorData.Println("Recover:", err)
	}
}
