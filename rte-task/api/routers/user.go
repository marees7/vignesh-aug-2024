package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, service *service.UserService) {
	user := &handler.UserHan{UserServices: service.User}

	r := router.Group("/user")
	{
		r.GET("users/allposts", user.GetAllJobPosts)
		r.GET("jobs/:job_title", user.GetJobByRole)
		r.GET("usercountry/:country", user.GetJobByCountry)
		r.GET("usersowndetails/user/:user_id", user.HandlerGetJobAppliedDetailsByUserId)
		r.POST("post/:user_id", user.ApplyJob)

	}
}
