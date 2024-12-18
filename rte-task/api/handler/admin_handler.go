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
	err := validation.ValidationJobPost(jobCreation, userID, roleType)
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

	c.JSON(http.StatusCreated, models.Response{
		Message: "Sucessfully Created the Details",
		Data:    jobCreation})
	loggers.InfoData.Println("Sucessfully Created the Details")
}

// admin updates their own JobPost
func (handler AdminHandler) UpdateJobPost(c *gin.Context) {
	var jobCreation models.JobCreation

	jobIDStr := c.Param("job_id")
	jobID, err := helpers.StringConvertion(jobIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		return
	}

	if err := c.ShouldBindJSON(&jobCreation); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured while request payload")
		return
	}
	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	//Validate each fields
	err = validation.ValidationUpdatePost(jobCreation, roleType, userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//update their job post by their IDs
	errorResponse := handler.Service.UpdateJobPost(&jobCreation, jobID)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error,
		})
		loggers.ErrorData.Println("Failed to update post")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":   "Post updated successfully",
		"JobStatus": jobCreation.JobStatus,
		"Vacancy":   jobCreation.Vacancy,
	})
	loggers.InfoData.Println("Sucessfully Updated the Details")
}

// admin get their jobs by Role
func (handler AdminHandler) GetPostByRole(c *gin.Context) {
	jobRole := c.Param("job_role")
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

	//get their jobDetailsBy role
	applicantDetails, errorResponse := handler.Service.GetPostByRole(jobRole, userID)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Can't able to get the Details Properly")
		return
	}
	c.JSON(http.StatusOK, models.Response{
		Message: "Successfully fetched the details",
		Data:    applicantDetails,
	})
	loggers.InfoData.Println("Successfully fetched job details")
}

func (handler AdminHandler) GetPostByJobID(c *gin.Context) {
	jobIDStr := c.Param("job_id")
	jobID, err := helpers.StringConvertion(jobIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		return
	}

	roleType := c.GetString("role_type")
	userID := c.GetInt("user_id")

	//Check their roles by admin or users
	if err := validation.ValidateRoleType(roleType); err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Validation failed:", err)
		return
	}

	// Call the service method
	applicantDetails, errorResponse := handler.Service.GetPostByJobID(jobID, userID)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, &models.Response{
			Error: errorResponse.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "Successfully fetched the details",
		Data:    applicantDetails,
	})
}

// admin get their jobs by UserId
func (handler AdminHandler) GetPostByUserID(c *gin.Context) {
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

	//Check their roles by admin or users
	if err := validation.ValidateRoleType(roleType); err != nil {
		c.JSON(http.StatusForbidden, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Validation failed:", err)
		return
	}

	//get their User's particular jobs By their userID's
	userJobDetails, errorResponse := handler.Service.GetPostByUserID(applicantID, userID)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error(),
		})
		loggers.ErrorData.Println("Failed to fetch job applied details:", err)
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "Successfully fetched job details",
		Data:    userJobDetails,
	})
	loggers.InfoData.Println("Successfully fetched job details")
}

// admin get their own posts by Admin
func (handler AdminHandler) GetPosts(c *gin.Context) {
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

	//get thier own post details By admin
	jobCreation, errorResponse := handler.Service.GetPosts(userID)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured while getting values by admin")
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully Get the details",
		Data:    jobCreation,
	})
	loggers.InfoData.Println("Sucessfully Get the Details by Admin")
}
