package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, service *service.UserService) {
	auth := &handler.AuthConnect{AuthService: service.Auth}
	r := router.Group("/auth")
	{
		//sign up their details
		r.POST("/signup", auth.SignUp)

		// Login with their Details
		r.POST("/login", auth.Login)
	}
}
