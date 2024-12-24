package config

import (
	"path/filepath"

	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/joho/godotenv"
)

func recoverPanic() {
	if r := recover(); r != nil {
		loggers.WarnData.Println("recovered from ", r)
	}
}

func LoadEnv() {
	defer recoverPanic()

	if err := godotenv.Load(filepath.Join(".env")); err != nil {
		panic(err)
	}
}
