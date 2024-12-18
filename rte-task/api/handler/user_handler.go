package handler

import (
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.IUserService
}

// user or admin get all job details
func (handler UserHandler) GetAllJobPosts(c *gin.Context) {
	country := c.Query("company_name")
	user, errorResponse := handler.Service.GetAllJobPosts(country)

	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error(),
		})
		loggers.ErrorData.Println("Sorry! I can't get your posts details properly")
		return
	}

	// Respond with the fetched data
	c.JSON(http.StatusOK, models.Response{
		Message: "Hurray! Successfully fetched the details",
		Data:    user,
	})
}

// user or admin get all jobrole and country
func (handler UserHandler) GetJobByRole(c *gin.Context) {
	jobTitle := c.Param("job_title")
	jobCountry := c.Param("country")

	//get thier job details by their JobRole
	jobCreation, errorResponse := handler.Service.GetJobByRole(jobTitle, jobCountry)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully Get the details by their JobRole",
		Data:    jobCreation,
	})
	loggers.InfoData.Println("Sucessfully Get the JobDetails By thier roles")
}

// user apply the job in that posts
func (handler UserHandler) CreateApplication(c *gin.Context) {
	var userJobDetails models.UserJobDetails

	if err := c.ShouldBindJSON(&userJobDetails); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to Invalid request payload")
		return
	}

	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	// Valid their User JobPost with Fields
	err := validation.ValidationUserJob(userJobDetails, roleType, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}

	// Check if user ID is applied for the Job or Not
	errorResponse := handler.Service.GetUserID(&userJobDetails)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		return
	}

	// user apply the job in that posts
	errorResponse = handler.Service.CreateApplication(&userJobDetails)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Message: "Sucessfully Applied Job ",
		Data:    userJobDetails})
	loggers.InfoData.Println("Sucessfully Applied the Job")
}

// user get by their userowndetails
func (handler UserHandler) GetUserAppliedJobs(c *gin.Context) {
	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	if err := validation.ValidateUserType(roleType); err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	//get their Details by userIds
	userJobDetails, errorResponse := handler.Service.GetUserAppliedJobs(userID)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully Get the details by thier own userIds",
		Data:    userJobDetails,
	})
	loggers.InfoData.Println("Sucessfully Get the details by thier own userIds")
}
