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
	r := router.Group("/v1/admin")
	{
		r.Use(middleware.Authenticate())

		//admin creates new JobPost
		r.POST("", admin.CreateJobPost)

		//get by jobrole and jobID
		r.GET("jobs", admin.GetApplicantAndJobDetails)

		//admin get by userid
		r.GET("/:user_id", admin.GetJobsAppliedByUser)

		//admin get by his id to know about how many post created
		r.GET("posts", admin.GetJobsCreated)

		//admins update by jobid and amin id
		r.PUT("/update/:job_id", admin.UpdateJobPost)

		//automaticaly delete jobPost
		r.DELETE("/delete", admin.DeleteJobPost)
	}
}
