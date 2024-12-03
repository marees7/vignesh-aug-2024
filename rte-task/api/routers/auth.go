package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, service *service.UserService) {
	user := &handler.AuthConnect{AuthService: service.Auth}
	r := router.Group("/auth")
	{
		r.POST("users/signup", user.SignUp)
		r.POST("users/login", user.Login)
	}

}
