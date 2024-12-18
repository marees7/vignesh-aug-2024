package config

import (
	"os"
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

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	if err = godotenv.Load(filepath.Join(filepath.Dir(wd), ".env")); err != nil {
		panic(err)
	}
}
