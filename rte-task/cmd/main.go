package main

import (
	"net/http"
	"os"

	"github.com/Vigneshwartt/golang-rte-task/api/routers"
	_ "github.com/Vigneshwartt/golang-rte-task/docs"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/internals/config"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Documenting API (JOB SITE)
// @version 1
// @contact.name Vigneshwartt
// @contact.url https://github.com/marees7/vignesh-aug-2024
// @contact.email vigneshwart2002@gmail.com
// @host localhost:8080

func init() {
	config.LoadEnv()
	loggers.LoggerFiles()
}

func main() {
	//Connect the Dbs
	dbconnection := internals.ConnectingDatabase()

	//migrate the tables
	dbconnection.Automigration()

	//send the service to handlers
	newrouter := gin.Default()

	routers.AuthRoutes(newrouter, dbconnection)
	routers.AdminRoutes(newrouter, dbconnection)
	routers.UserRoutes(newrouter, dbconnection)

	newrouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	//start the server
	loggers.InfoData.Println("Server started on port")
	err := http.ListenAndServe(os.Getenv("HTTP_PORT"), newrouter)
	if err != nil {
		loggers.ErrorData.Fatalln("Failed to start the server", err)
	}
}
