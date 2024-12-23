package internals

import (
	"fmt"
	"os"

	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type NewConnection struct {
	*gorm.DB
}

func ConnectingDatabase() *NewConnection {
	path := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PORT"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	Connection, err := gorm.Open(postgres.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	
	defer HandlePanic()
	loggers.InfoData.Println("Connected sucessfully")
	return &NewConnection{
		Connection,
	}
}

func HandlePanic() {
	if err := recover(); err != nil {
		loggers.ErrorData.Println("Recover:", err)
	}
}
