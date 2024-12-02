package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/handler"
	"github.com/Vigneshwartt/golang-rte-task/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userhands handler.UserHandler) {
	r := router.Group("/user")
	{
		r.POST("insert/jobs", middleware.Authenticate(), userhands.CreateJobPost)
		r.GET("users/allposts", userhands.GetAllJobPosts)
		r.GET("jobs/:job_title", userhands.GetJobByRole)
		r.GET("usercountry/:country", userhands.GetJobByCountry)
	}
}
