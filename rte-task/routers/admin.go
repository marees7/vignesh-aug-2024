package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/handler"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine, userhands handler.UserHandler) {
	r := router.Group("/admin")
	{
		r.POST("users/signup", userhands.SignUp)
		r.POST("users/login", userhands.Login)
	}
}
