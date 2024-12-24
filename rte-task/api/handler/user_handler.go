package handler

import (
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/common/helpers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	Service service.IUserService
}

// @Summary CreateApplication
// @Description User apply the job in that posts
// @Security ApiKeyAuth
// @Name Type Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer)
// @Param user body models.UserJobDetails true "UserJobs"
// @Tags User
// @Accept json
// @Produce json
// @Success 201 {object}  dto.Response
// @Failure 500 {object} dto.Response
// @Failure 422 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /v1/user [post]
func (handler UserHandler) CreateApplication(c *gin.Context) {
	var userJobDetails models.UserJobDetails

	if err := c.ShouldBindJSON(&userJobDetails); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to Invalid request payload-", err)
		return
	}

	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	// Valid their User JobPost with Fields
	err := validation.ValidateUserApplicaton(userJobDetails, roleType, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error Occured validating user application-", err)
		return
	}

	// Check if user ID is applied for the Job or Not
	errorResponse := handler.Service.GetUserByID(&userJobDetails)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, dto.Response{
			Error: errorResponse.Error.Error()})
		return
	}

	// user apply the job in that posts
	errorResponse = handler.Service.CreateApplication(&userJobDetails)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, dto.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured while creating values-", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Sucessfully Applied the Job-", userJobDetails.JobID)
	c.JSON(http.StatusCreated, dto.Response{
		Message: "Sucessfully Applied Job ",
		Data:    userJobDetails})
}

// @Summary Get All Job Posts
// @Description User or admin get all job details
// @Param company_name query string false "Search by company_name"
// @Param job_role query string false "Search by job_role"
// @Param country query string false "Search by country"
// @Param limit query int false "Search by limit"
// @Param offset query int false "Search by offset"
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}  dto.Response
// @Failure 500 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Router /posts [get]
func (handler UserHandler) GetAllJobPosts(c *gin.Context) {
	companyName := c.Query("company_name")
	jobRole := c.Query("job_role")
	jobCountry := c.Query("country")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	limit, offset, errorResponse := helpers.Pagination(offsetStr, limitStr)
	if errorResponse != nil {
		loggers.ErrorData.Println("Error Occured when fetching paganition-", errorResponse.Error)
	}

	searchJobs := map[string]interface{}{
		"company": companyName,
		"role":    jobRole,
		"country": jobCountry,
		"limit":   limit,
		"offset":  offset,
	}

	jobCreation, errorResponse, count := handler.Service.GetAllJobPosts(searchJobs)

	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, dto.Response{
			Error: errorResponse.Error.Error(),
		})
		loggers.ErrorData.Println("Sorry! I can't get your posts details properly-", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Sucessfully fetched the JobPosts")
	// Respond with the fetched data
	c.JSON(http.StatusOK, dto.Response{
		Message: "Hurray! Successfully fetched the JobPosts",
		Data:    jobCreation,
		Total:   count,
		Limit:   limit,
		Offset:  offset,
	})
}

// @Summary GetUserAppliedJobs
// @Description User get by their userowndetails
// @Security ApiKeyAuth
// @Name Type Bearer
// @Param Authorization header string true "Insert your access token" default(Bearer)
// @Param limit query int false "Search by limit"
// @Param offset query int false "Search by offset"
// @Tags User
// @Accept json
// @Produce json
// @Success 200 {object}  dto.Response
// @Failure 403 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 500 {object} dto.Response
// @Router /v1/user [get]
func (handler UserHandler) GetUserAppliedJobs(c *gin.Context) {
	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")

	if err := validation.ValidateUserType(roleType); err != nil {
		c.JSON(http.StatusForbidden, dto.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values-", err)
		return
	}

	limit, offset, errorResponse := helpers.Pagination(offsetStr, limitStr)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, dto.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error Occured", errorResponse.Error)
		return
	}

	userJobs := map[string]interface{}{
		"userID": userID,
		"limit":  limit,
		"offset": offset,
	}

	//get their Details by userIds
	userJobDetails, errorResponse, count := handler.Service.GetUserAppliedJobs(userJobs)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, dto.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured while getting values-", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Sucessfully Get the details by thier ID-", userID)
	c.JSON(http.StatusOK, dto.Response{
		Message: "Sucessfully Get the details by thier ID",
		Data:    userJobDetails,
		Total:   count,
		Limit:   limit,
		Offset:  offset,
	})
}
