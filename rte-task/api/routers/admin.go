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
	adminRepos := repository.GetAdminRepository(dbconnection)

	// send the repos to service
	adminservice := service.GetAdminService(adminRepos)

	//send service to handler
	admin := &handler.AdminNewHandler{I_AdminService: adminservice}
	r := router.Group("/admin")
	{
		r.Use(middleware.Authenticate())
		//admin creates new JobPost
		r.POST("insert", admin.CreateJobPost)

		//admin get by jobid userid
		r.GET("userjobsbyid/:job_id", admin.GetPostByJobID)

		//admin get by userid
		r.GET("userid/:user_id", admin.GetPostByUserID)

		//get by role and userid
		r.GET("userdetails/:job_role", admin.GetPostByRole)

		//admins update by jobid and amin id
		r.PUT("/update/:job_id", admin.UpdatePost)

		//admin get by his id to know about how many psot created
		r.GET("postdetails", admin.GetOwnPosts)
	}
}
