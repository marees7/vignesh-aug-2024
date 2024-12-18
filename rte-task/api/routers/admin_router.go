package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/middleware"
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/gin-gonic/gin"
)

func AdminRoutes(router *gin.Engine, dbconnection *internals.ConnectionNew) {
	//send the Db connection to repos
	adminRepos := repository.InitAdminRepo(dbconnection)

	// send the repos to service
	adminservice := service.InitAdminService(adminRepos)

	//send service to handler
	admin := &handler.AdminHandler{Service: adminservice}
	r := router.Group("/admin")
	{
		r.Use(middleware.Authenticate())
		//admin creates new JobPost
		r.POST("insert", admin.CreateJobPost)

		//admins update by jobid and amin id
		r.PUT("/update/:job_id", admin.UpdateJobPost)

		//admin get by jobid userid
		r.GET("userjobsbyid/:job_id", admin.GetPostByJobID)

		//get by role and userid
		r.GET("userdetails/:job_role", admin.GetPostByRole)

		//admin get by userid
		r.GET("userid/:user_id", admin.GetPostByUserID)

		//admin get by his id to know about how many psot created
		r.GET("postdetails", admin.GetPosts)
	}
}
