package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/common/helpers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type UserHan struct {
	service.UserServices
}

func (database UserHan) GetAllJobPosts(c *gin.Context) {
	var user []models.JobCreation
	userType := c.GetString("role_type")

	dataOfJobs := database.ServiceGetAllPostDetails(&user, userType)
	if dataOfJobs != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: dataOfJobs.Error})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}
	for _, values := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details",
			Data:    values,
		})
	}
	loggers.InfoData.Println("Sucessfuly Get the all created Post Details")
}

func (database UserHan) GetJobByRole(c *gin.Context) {
	var user []models.JobCreation
	jobs := c.Param("job_title")
	country := c.Param("country")

	fmt.Println("jobs", jobs)
	userType := c.GetString("role_type")

	dbValues := database.ServiceGetJobDetailsByRole(&user, jobs, country, userType)
	if dbValues != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: "Could not able to get the details"})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}
	for _, values := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details by their JobRole",
			Data:    values,
		})
	}
	loggers.InfoData.Println("Sucessfully Get the JobDetails By thier roles")
}

func (database UserHan) ApplyJob(c *gin.Context) {
	var user models.UserJobDetails

	value := c.Param("user_id")
	applyuserid, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}

	fmt.Println("values", applyuserid)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to Invalid request payload")
		return
	}
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	fmt.Println("userid", userid)
	fmt.Println("usertype", userType)
	Dbvalues := database.ApplyJobPost(&user, userType, userid, applyuserid)
	if Dbvalues != nil {
		c.JSON(500, models.CommonResponse{
			Error: "Error occured while creating values"})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}
	c.JSON(http.StatusOK, models.CommonResponse{
		Message: "Sucessfully Applied Job ",
		Data:    user})
	loggers.InfoData.Println("Sucessfully Applied the Job")
}

func (database UserHan) HandlerGetJobAppliedDetailsByUserId(c *gin.Context) {
	var user []models.UserJobDetails

	if err := helpers.CheckuserType(c, "USER"); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	value := c.Param("user_id")
	values, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")
	fmt.Println("values", values)
	fmt.Println("userid", userid)
	fmt.Println("usertype", userType)

	dbvalues := database.GetJobAppliedDetailsByUserId(&user, values, userid)
	if dbvalues != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: "OOPS! ID is mismatching here,Enter properly"})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	for _, valuess := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details by thier own userIds",
			Data:    valuess,
		})
	}
	loggers.InfoData.Println("Sucessfully Get the details by thier own userIds")
}
