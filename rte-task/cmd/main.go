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

func main() {
	dbconnection := internals.ConnectingDatabase()

	adminRepos := repository.GetAdminRepository(dbconnection)
	authrepo := repository.GetAuthRepository(dbconnection)
	userrepo := repository.GetUserRepository(dbconnection)

	adminservice := service.GetAdminService(adminRepos)
	authservice := service.GetAuthService(authrepo)
	userservice := service.GetUserService(userrepo)

	newrouter := gin.Default()
	routers.AuthRoutes(newrouter, authservice)
	routers.UserRoutes(newrouter, userservice)
	routers.AdminRouter(newrouter, adminservice)

	log.Println("Server started on port :8080")
	http.ListenAndServe(":8080", newrouter)
}
