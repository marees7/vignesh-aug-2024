package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/handler"
	"github.com/Vigneshwartt/golang-rte-task/middleware"
	"github.com/gin-gonic/gin"
)

func IntializeRouter(userhands handler.UserHandler) *gin.Engine {
	r := gin.Default()

	r.POST("users/signup", userhands.SignUp)
	r.POST("users/login", userhands.Login)
	r.GET("users", middleware.Authenticate(), userhands.GetUserss)
	r.GET("users/:role_id", middleware.Authenticate(), userhands.GetUser)
	r.POST("users/post", userhands.CreateJobPost)
	return r
}
