package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/middleware"
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, dbconnection *internals.ConnectionNew) {
	//send the Db connection to repos
	userrepo := repository.InitUserRepo(dbconnection)

	// send the repos to service
	userservice := service.InitUserService(userrepo)

	//send service to handler
	user := handler.UserHandler{Service: userservice}

	r := router.Group("/user")
	{

		//user apply the job in that posts
		r.POST("create", middleware.Authenticate(), user.CreateApplication)

		//user or admin get all job details
		r.GET("allposts", user.GetAllJobPosts)

		//user or admin get all jobrole and country
		r.GET("jobs/:job_title/:country", user.GetJobByRole)

		//user get by their userowndetails
		r.GET("userid", middleware.Authenticate(), user.GetUserAppliedJobs)
	}
}
