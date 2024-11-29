package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/handler"
	"github.com/Vigneshwartt/golang-rte-task/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, userhands handler.UserHandler) {
	r := router.Group("/user")
	{
		// r.Use(middleware.Authenticate())
		r.GET("users/allposts", userhands.GetJobPost)
		r.GET("jobs/:job_title", userhands.GetJobByRole)
		r.GET("usercountry/:country", userhands.GetJobByCountry)
		r.POST("insert/jobs", middleware.Authenticate(), userhands.CreateJobPost)
	}
}
