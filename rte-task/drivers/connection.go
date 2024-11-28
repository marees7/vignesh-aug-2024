package drivers

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectingDatabase() *gorm.DB {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error failed to load the env file ")
		return nil
	}
	host := os.Getenv("DB_host")
	user := os.Getenv("DB_user")
	port := os.Getenv("DB_port")
	password := os.Getenv("DB_password")
	dbname := os.Getenv("DB_dbname")
	path := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s", host, user, port, password, dbname)
	Connection, err := gorm.Open(postgres.Open(path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	GlobalConnection = Connection
	Automigration()
	fmt.Println("Connection make sucessfully")
	return Connection
}

func HandlePanic() {
	if err := recover(); err != nil {
		fmt.Println("Recover:", err)
	}
}
