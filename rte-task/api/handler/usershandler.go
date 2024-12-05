package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
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

	err := database.ServiceGetAllPostDetails(&user, userType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Sorry! I can't get your all posts details properly")
		return
	}
	for _, values := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Hurray!,Sucessfully Get the details",
			Data:    values,
		})
	}
	loggers.InfoData.Println("Sucessfuly Get the all created Post Details")
}

func (database UserHan) GetJobByRole(c *gin.Context) {
	var user []models.JobCreation
	jobs := c.Param("job_title")
	country := c.Param("country")

	// fmt.Println("jobs", jobs)
	userType := c.GetString("role_type")

	err := database.ServiceGetJobDetailsByRole(&user, jobs, country, userType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
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
	var newpost models.JobCreation

	value := c.Param("user_id")
	paramid, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}

	// fmt.Println("values", paramid)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to Invalid request payload")
		return
	}

	tokentype := c.GetString("role_type")
	tokenid := c.GetInt("user_id")

	// fmt.Println("userid", tokenid)
	// fmt.Println("usertype", tokentype)
	err = validation.ValidationUserJob(user, tokentype, tokenid, paramid)
	if err != nil {
		c.JSON(500, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}

	err = database.CheckJobId(&user, &newpost)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		return
	}

	err = database.ApplyJobPost(&user)
	if err != nil {
		c.JSON(500, models.CommonResponse{
			Error: err.Error()})
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
	userid := c.GetInt("user_id")
	fmt.Println("values", values)
	fmt.Println("userid", userid)

	err = database.GetJobAppliedDetailsByUserId(&user, values, userid)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
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
