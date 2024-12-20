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
		loggers.ErrorData.Println("Cant able to get the JobPost Details")
		return
	}

	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	//validate each fields in Jobpost
	err := validation.ValidateJobPost(jobCreation, userID, roleType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//create their Jobposts
	errorResponse := handler.Service.CreateJobPost(&jobCreation)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse})
		loggers.ErrorData.Println("OOPS! Your Id or roletype is Mismatching Here,Check It Once Again")
		return
	}

	loggers.InfoData.Println("Sucessfully Created the Details")
	c.JSON(http.StatusCreated, models.Response{
		Message: "Sucessfully Created the Details",
		Data:    jobCreation})
}

// admin get their own posts by Admin
func (handler AdminHandler) GetJobsCreated(c *gin.Context) {
	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	limitStr := c.Query("Limit")
	offsetStr := c.Query("Offset")

	limit, offset := helpers.Pagination(offsetStr, limitStr)

	createdPosts := map[string]int{
		"Limit":  limit,
		"Offset": offset,
		"UserID": userID,
	}

	// Check their roles by admin or users
	err := validation.ValidateRoleType(roleType)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//get thier own post details By admin
	jobCreation, errorResponse, count := handler.Service.GetJobsCreated(createdPosts)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Failed to fetch job created details")
		return
	}

	loggers.InfoData.Println("Sucessfully fetched the job Posted by Admin")
	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully fetched the job posted by Admin",
		Data:    jobCreation,
		Limit:   limit,
		Offset:  offset,
		Total:   count,
	})
}

// admin get their jobs by Role
func (handler AdminHandler) GetApplicantAndJobDetails(c *gin.Context) {
	jobRole := c.Query("JobRole")
	jobIDStr := c.Query("JobID")
	offsetStr := c.Query("Offset")
	limitStr := c.Query("Limit")

	limit, offset := helpers.Pagination(offsetStr, limitStr)

	jobID, _ := strconv.Atoi(jobIDStr)
	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	// Check their roles by admin or users
	err := validation.ValidateRoleType(roleType)
	if err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	jobDetailsMap := map[string]interface{}{
		"JobRole": jobRole,
		"JobID":   jobID,
		"UserID":  userID,
		"Limit":   limit,
		"Offset":  offset,
	}

	//get their jobDetailsBy role
	applicantDetails, errorResponse, count := handler.Service.GetApplicantAndJobDetails(jobDetailsMap)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Can't able to get the Details Properly")
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
	limitStr := c.Query("Limit")
	offsetStr := c.Query("Offset")
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
	limit, offset := helpers.Pagination(offsetStr, limitStr)

	//Check their roles by admin or users
	if err := validation.ValidateRoleType(roleType); err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Validation failed:", err)
		return
	}

	userIDJobs := map[string]interface{}{
		"Limit":     limit,
		"Offset":    offset,
		"Applicant": applicantID,
		"UserID":    userID,
	}

	//get their User's particular jobs By their userID's
	userJobDetails, errorResponse, count := handler.Service.GetJobsAppliedByUser(userIDJobs)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error(),
		})
		loggers.ErrorData.Println("Failed to fetch job applied details")
		return
	}

	loggers.InfoData.Println("Successfully fetched User applied Jobs")
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
	var jobUpdation models.JobCreation

	jobIDStr := c.Param("job_id")
	jobID, err := helpers.StringConvertion(jobIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&jobUpdation); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured while request payload")
		return
	}

	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	//Validate each fields
	err = validation.ValidateUpdatePost(jobUpdation, roleType, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//update their job post by their IDs
	errorResponse := handler.Service.UpdateJobPost(&jobUpdation, jobID, userID)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error,
		})
		loggers.ErrorData.Println("Failed to update post")
		return
	}

	loggers.InfoData.Println("Sucessfully Updated the JobPost")
	c.JSON(http.StatusOK, models.Response{
		Message:   "Sucessfully Updated the JobPost",
		JobStatus: jobUpdation.JobStatus,
		Vacancy:   jobUpdation.Vacancy,
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
			Error: errorResponse.Error,
		})
		loggers.ErrorData.Println("Failed to Delete post")
		return
	}

	loggers.InfoData.Println("Sucessfully Deleted the JobPost")
	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully Deleted the JobPost",
	})
}
