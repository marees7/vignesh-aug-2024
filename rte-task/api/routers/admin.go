package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/middleware"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine, service *service.UserService) {
	admin := &handler.AdminHand{AdminService: service.Admin}
	r := router.Group("/admin")
	{
		r.Use(middleware.Authenticate())
		//admin creates new JobPost
		r.POST("insert/:admin_id", admin.CreateJobPost)

		//admin get by jobid userid
		r.GET("userjobsbyid/:job_id/:admin_id", admin.GetJobAppliedDetailsByJobId)

		//admin get by userid
		r.GET("userid/:user_id/:admin_id", admin.GetJobAppliedDetailsByUserId)

		//get by role and userid
		r.GET("userdetails/:job_role/:admin_id", admin.GetJobAppliedDetailsbyrole)

		//admins update by jobid and amin id
		r.PUT("/update/:job_id/:admin_id", admin.UpdatePost)

		//admin get by his id to know about how many psot created
		r.GET("postdetails/:admin_id", admin.GetJobsByAdmin)
	}
}
