package main

import (
	"net/http"
	"os"

	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/api/routers"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/gin-gonic/gin"
)

func init() {
	//first Load the db connection
	internals.ConnectingDatabase()
	//migrate the tables
	internals.Automigration()
}
func main() {
	//Connect the Db
	dbconnection := internals.GetConnection()

	//send the Db connection to repos
	adminRepos := repository.GetAdminRepository(dbconnection)
	authrepo := repository.GetAuthRepository(dbconnection)
	userrepo := repository.GetUserRepository(dbconnection)

	// send the repos to service
	adminservice := service.GetAdminService(adminRepos)
	authservice := service.GetAuthService(authrepo)
	userservice := service.GetUserService(userrepo)

	//send the service to handlers
	newrouter := gin.Default()
	routers.AuthRoutes(newrouter, authservice)
	routers.UserRoutes(newrouter, userservice)
	routers.AdminRouter(newrouter, adminservice)

	//start the server
	loggers.InfoData.Println("Server started on port")
	err := http.ListenAndServe(os.Getenv("HTTP_PORT"), newrouter)
	if err != nil {
		loggers.ErrorData.Fatalln("Failed to start the server", err)
	}
}
