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
		r.GET("users", userhands.MultipleUsers)
		r.GET("users/:role_id", userhands.GetUser)
		r.POST("users/post", userhands.ApplyJob)

		r.GET("admin/alljobs", userhands.GetAllAppliedJobDetails)
		r.GET("userjobsbyid/admin/:job_id", userhands.GetJobAppliedDetailsByJobId)
		r.GET("userid/admin/:user_id", userhands.GetJobAppliedDetailsByUserId)
		r.GET("userdetails/admin/:job_role", userhands.GetJobAppliedDetailsbyrole)
		r.GET("usersowndetails/user/:user_id", userhands.HandlerGetJobAppliedDetailsByUserId)
		r.PUT("/update/:job_id_new", userhands.UpdatePost)
		// r.DELETE("/delete/:job_id_new", userhands.DeletePost)
	}
}
