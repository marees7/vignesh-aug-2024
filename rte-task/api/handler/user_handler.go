package handler

import (
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/common/helpers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.IUserService
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
	err := validation.ValidateUserApplicaton(userJobDetails, roleType, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}

	// Check if user ID is applied for the Job or Not
	errorResponse := handler.Service.GetUserByID(&userJobDetails)
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
	loggers.InfoData.Println("Sucessfully Applied the Job")

	c.JSON(http.StatusCreated, models.Response{
		Message: "Sucessfully Applied Job ",
		Data:    userJobDetails})
}

// user or admin get all job details
func (handler UserHandler) GetAllJobPosts(c *gin.Context) {
	companyName := c.Query("CompanyName")
	jobRole := c.Query("JobRole")
	jobCountry := c.Query("Country")
	limitStr := c.Query("Limit")
	offsetStr := c.Query("Offset")

	limit, offset := helpers.Pagination(offsetStr, limitStr)

	searchJobs := map[string]interface{}{
		"Company": companyName,
		"Role":    jobRole,
		"Country": jobCountry,
		"Limit":   limit,
		"Offset":  offset,
	}

	jobCreation, errorResponse, count := handler.Service.GetAllJobPosts(searchJobs)

	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error(),
		})
		loggers.ErrorData.Println("Sorry! I can't get your posts details properly")
		return
	}
	loggers.InfoData.Println("Sucessfully fetched the JobPosts ")

	// Respond with the fetched data
	c.JSON(http.StatusOK, models.Response{
		Message: "Hurray! Successfully fetched the JobPosts",
		Data:    jobCreation,
		Total:   count,
		Limit:   limit,
		Offset:  offset,
	})
}

// user get by their userowndetails
func (handler UserHandler) GetUserAppliedJobs(c *gin.Context) {
	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")
	limitStr := c.Query("Limit")
	offsetStr := c.Query("Offset")

	if err := validation.ValidateUserType(roleType); err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}
	limit, offset := helpers.Pagination(offsetStr, limitStr)

	userJobs := map[string]interface{}{
		"UserID": userID,
		"Limit":  limit,
		"Offset": offset,
	}

	//get their Details by userIds
	userJobDetails, errorResponse, count := handler.Service.GetUserAppliedJobs(userJobs)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}
	loggers.InfoData.Println("Sucessfully Get the details by thier ID")

	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully Get the details by thier ID",
		Data:    userJobDetails,
		Total:   count,
		Limit:   limit,
		Offset:  offset,
	})
}
