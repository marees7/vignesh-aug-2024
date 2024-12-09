package internals

import (
	"fmt"
	"os"

	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	GlobalConnection *gorm.DB
)

func ConnectingDatabase() {
	err := godotenv.Load(".env")
	if err != nil {
		loggers.ErrorData.Println("Error failed to load the env file ")
		return
	}

	path := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s", os.Getenv("DB_host"), os.Getenv("DB_user"), os.Getenv("DB_port"), os.Getenv("DB_password"), os.Getenv("DB_dbname"))
	Connection, err := gorm.Open(postgres.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	GlobalConnection = Connection
	defer HandlePanic()
	loggers.InfoData.Println("Connected sucessfully")
}

func GetConnection() *gorm.DB {
	return GlobalConnection
}

func HandlePanic() {
	if err := recover(); err != nil {
		loggers.ErrorData.Println("Recover:", err)
	}
}
