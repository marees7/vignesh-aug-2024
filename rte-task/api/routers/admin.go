package routers

import (
	"github.com/Vigneshwartt/golang-rte-task/api/handler"
	"github.com/Vigneshwartt/golang-rte-task/api/middleware"
	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/gin-gonic/gin"
)

func AdminRouter(router *gin.Engine, service *service.UserService) {
	user := &handler.AdminHand{AdminService: service.Admin}
	r := router.Group("/admin")
	{
		r.Use(middleware.Authenticate())
		// r.GET("users", user.MultipleUsers)
		r.POST("insert/jobs/:user_id", middleware.Authenticate(), user.CreateJobPost)
		// r.GET("users/:user_id/:admin_id", userhands.GetUser)

		// r.GET("admin/alljobs/:user_id", user.GetAllAppliedJobDetails)

		//admin get by jobid userid--yes
		r.GET("userjobsbyid/admin/:job_id/:user_id", user.GetJobAppliedDetailsByJobId)
		//admin get by userid --yes
		r.GET("userid/admin/:user_id/:admin_id", user.GetJobAppliedDetailsByUserId)
		//get by role and userid
		r.GET("userdetails/admin/:job_role/:user_id", user.GetJobAppliedDetailsbyrole)
		//users get by id

		r.PUT("/update/:job_id/:user_id", user.UpdatePost)
		// r.DELETE("/delete/:job_id", user.DeletePost)
	}
}
