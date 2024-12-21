package handler

import (
	"net/http"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/common/helpers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	Service service.IAdminService
}

// admin creates new JobPost
func (handler AdminHandler) CreateJobPost(c *gin.Context) {
	var jobCreation models.JobCreation

	if err := c.BindJSON(&jobCreation); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Cant able to get the JobPost Detail", err)
		return
	}

	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	//validate each fields in Jobpost
	err := validation.ValidateJobPost(jobCreation, userID, roleType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetail", err)
		return
	}

	//create their Jobposts
	errorResponse := handler.Service.CreateJobPost(&jobCreation)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println(" Error occured,", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Sucessfully Created the JobPost Detail", jobCreation.JobID)
	c.JSON(http.StatusCreated, models.Response{
		Message: "Sucessfully Created the JobPost Detail",
		Data:    jobCreation})
}

// admin get their own posts by Admin
func (handler AdminHandler) GetJobsCreated(c *gin.Context) {
	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	jobRole := c.Query("job_role")
	country := c.Query("country")

	limit, offset, errorResponse := helpers.Pagination(offsetStr, limitStr)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error Occured", errorResponse.Error)
		return
	}

	searchMap := map[string]interface{}{
		"limit":   limit,
		"offset":  offset,
		"userID":  userID,
		"jobRole": jobRole,
		"country": country,
	}

	// Check their roles by admin or users
	err := validation.ValidateRoleType(roleType)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails", err)
		return
	}

	//get thier own post details By admin
	jobCreation, errorResponse, count := handler.Service.GetJobsCreated(searchMap)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Failed to fetch job created details", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Sucessfully fetched the JobCreated Details", userID)
	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully fetched the Jobcreated Details",
		Data:    jobCreation,
		Limit:   limit,
		Offset:  offset,
		Total:   count,
	})
}

// admin get their jobs by Role
func (handler AdminHandler) GetApplicantAndJobDetails(c *gin.Context) {
	jobRole := c.Query("job_role")
	jobIDStr := c.Query("job_id")
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")

	limit, offset, errorResponse := helpers.Pagination(offsetStr, limitStr)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error Occured", errorResponse.Error)
		return
	}

	jobID, _ := strconv.Atoi(jobIDStr)
	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	// Check their roles by admin or users
	err := validation.ValidateRoleType(roleType)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails", err)
		return
	}

	jobDetailsMap := map[string]interface{}{
		"jobRole": jobRole,
		"jobID":   jobID,
		"userID":  userID,
		"limit":   limit,
		"offset":  offset,
	}

	//get their jobDetailsBy role
	applicantDetails, errorResponse, count := handler.Service.GetApplicantAndJobDetails(jobDetailsMap)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Can't able to get the Details Properly", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Successfully fetched Job Posts ")
	c.JSON(http.StatusOK, models.Response{
		Message: "Successfully fetched Job Posts ",
		Data:    applicantDetails,
		Total:   count,
		Limit:   limit,
		Offset:  offset,
	})
}

// admin get their jobs by UserId
func (handler AdminHandler) GetJobsAppliedByUser(c *gin.Context) {
	limitStr := c.Query("limit")
	offsetStr := c.Query("offset")
	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	applicantIDStr := c.Param("user_id")
	applicantID, err := helpers.StringConvertion(applicantIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		return
	}

	limit, offset, errorResponse := helpers.Pagination(offsetStr, limitStr)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error Occured", errorResponse.Error)
		return
	}

	//Check their roles by admin or users
	if err := validation.ValidateRoleType(roleType); err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Validation failed:", err)
		return
	}

	userIDJobs := map[string]interface{}{
		"limit":     limit,
		"offset":    offset,
		"applicant": applicantID,
		"userID":    userID,
	}

	//get their User's particular jobs By their userID's
	userJobDetails, errorResponse, count := handler.Service.GetJobsAppliedByUser(userIDJobs)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error(),
		})
		loggers.ErrorData.Println("Failed to fetch job applied details", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Successfully fetched User applied Jobs", userID)
	c.JSON(http.StatusOK, models.Response{
		Message: "Successfully fetched User applied Jobs",
		Data:    userJobDetails,
		Total:   count,
		Limit:   limit,
		Offset:  offset,
	})
}

// admin updates their own JobPost
func (handler AdminHandler) UpdateJobPost(c *gin.Context) {
	var jobData models.JobCreation

	jobIDStr := c.Param("job_id")
	jobID, err := helpers.StringConvertion(jobIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&jobData); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured while request payload", err)
		return
	}

	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	//Validate each fields
	err = validation.ValidateUpdatePost(jobData, roleType, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails", err)
		return
	}

	//update their job post by their IDs
	errorResponse := handler.Service.UpdateJobPost(&jobData, jobID, userID)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error(),
		})
		loggers.ErrorData.Println("Failed to update post", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Sucessfully Updated the JobPost")
	c.JSON(http.StatusOK, models.Response{
		Message:   "Sucessfully Updated the JobPost",
		JobStatus: jobData.JobStatus,
		Vacancy:   jobData.Vacancy,
	})
}

func (handler AdminHandler) DeleteJobPost(c *gin.Context) {
	roleType := c.GetString("role_type")

	//Check their roles by admin or users
	if err := validation.ValidateRoleType(roleType); err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Validation failed:", err)
		return
	}

	errorResponse := handler.Service.DeleteJobPost()
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error(),
		})
		loggers.ErrorData.Println("Failed to Delete post", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Sucessfully Deleted the JobPost")
	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully Deleted the JobPost",
	})
}
