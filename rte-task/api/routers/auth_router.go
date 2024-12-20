package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, dbconnection *internals.ConnectionNew) {
	//send the Db connection to repos
	authrepo := repository.InitAuthRepo(dbconnection)

	// send the repos to service
	authservice := service.InitAuthService(authrepo)

	//send service to handler
	auth := &handler.AuthHandler{Service: authservice}
	r := router.Group("/v1/auth")
	{
		//sign up their details
		r.POST("/signup", auth.CreateUser)

		// Login with their Details
		r.POST("/login", auth.GetUserDetail)
	}
}
