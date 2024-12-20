package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/middleware"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, service *service.UserService) {
	user := &handler.UserHan{UserServices: service.User}

	r := router.Group("/user")
	{
		r.Use(middleware.Authenticate())
		//user apply the job in that posts
		r.POST("post/:user_id", user.UsersApplyForJobs)

		//user or admin get all job details
		r.GET("users/allposts", user.GetAllJobPosts)

		//user or admin get all jobrole and country
		r.GET("jobs/:job_title/:country", user.GetJobByRole)

		//user or admin get companyName
		r.GET("company/:company_name", user.GetByCompanyname)

		//user get by their userowndetails
		r.GET("usersowndetails/user/:user_id", user.UsersGetTheirDetailsByTheirownIds)
	}
}
