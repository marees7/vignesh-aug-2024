package main

import (
	"net/http"
	"os"

	"github.com/Vigneshwartt/golang-rte-task/api/routers"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/internals/config"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/gin-gonic/gin"
)

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

	routers.AdminRoutes(newrouter, dbconnection)
	routers.AuthRoutes(newrouter, dbconnection)
	routers.UserRoutes(newrouter, dbconnection)

	//start the server
	loggers.InfoData.Println("Server started on port")
	err := http.ListenAndServe(os.Getenv("HTTP_PORT"), newrouter)
	if err != nil {
		loggers.ErrorData.Fatalln("Failed to start the server", err)
	}
}
