package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/handler"
	"github.com/Vigneshwartt/golang-rte-task/middleware"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine, userhands handler.UserHandler) {
	r := router.Group("/auth")
	{
		r.Use(middleware.Authenticate())
		r.GET("users", userhands.GetUserss)
		r.GET("users/:role_id", userhands.GetUser)
		r.POST("users/post", userhands.ApplyJob)
		r.GET("userdetails/admin/:job_role", userhands.GetJobAppliedDetails)
		r.GET("user/alljobs", userhands.GetAllJobDetails)
	}
}
