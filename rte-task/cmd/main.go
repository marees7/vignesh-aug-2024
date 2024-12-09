package main

import (
	"log"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/api/routers"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/internals"
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
	log.Println("Server started on port :8080")
	http.ListenAndServe(":8080", newrouter)
}
