package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/middleware"
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, dbconnection *internals.NewConnection) {
	//send the Db connection to repos
	userrepo := repository.InitUserRepo(dbconnection)

	// send the repos to service
	userservice := service.InitUserService(userrepo)

	//send service to handler
	user := handler.UserHandler{Service: userservice}

	// user or admin get all job details
	router.GET("posts", user.GetAllJobPosts)

	r := router.Group("/v1/user")
	{
		r.Use(middleware.Authenticate())

		//user apply the job in that posts
		r.POST("", user.CreateApplication)

		//user get by their userowndetails
		r.GET("", user.GetUserAppliedJobs)
	}
}
