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

type AdminNewHandler struct {
	service.I_AdminService
}

// admin creates new JobPost
func (service AdminNewHandler) CreateJobPost(c *gin.Context) {
	var user models.JobCreation

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Cant able to get the JobPost Details")
		return
	}
	roleType := c.GetString("role_type")
	roleID := c.GetInt("user_id")

	//validate each fields in Jobpost
	err := validation.ValidationJobPost(user, roleID, roleType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//create their Jobposts
	err = service.CreateJobPosts(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("OOPS! Your Id or roletype is Mismatching Here,Check It Once Again")
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Message: "Sucessfully Created the Details",
		Data:    user})
	loggers.InfoData.Println("Sucessfully Created the Details")
}

// admin updates their own JobPost
func (service AdminNewHandler) UpdatePost(c *gin.Context) {
	var post models.JobCreation

	paramID := c.Param("job_id")
	jobID, err := strconv.Atoi(paramID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: "Error occured while String Convertion,Please check properly",
		})
		loggers.ErrorData.Println("Error occured Invalid JobId Found,Please Check")
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured while request payload")
		return
	}
	roleType := c.GetString("role_type")
	roleID := c.GetInt("user_id")

	//Validate each fields
	err = validation.ValidationUpdatePost(post, roleType, roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//update their job post by their IDs
	if err := service.UpdatePosts(&post, jobID); err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to update post")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":   "Post updated successfully",
		"JobStatus": post.JobStatus,
		"Vacancy":   post.Vacancy,
	})
	loggers.InfoData.Println("Sucessfully Updated the Details")
}

// admin get their jobs by Role
func (service AdminNewHandler) GetPostByRole(c *gin.Context) {
	paramJobRole := c.Param("job_role")
	roleType := c.GetString("role_type")
	roleID := c.GetInt("user_id")

	// Check their roles by admin or users
	err := helpers.ValidateRoleType(roleType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//get their jobDetailsBy role
	user, err := service.GetAllPostsByRole(paramJobRole, roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Can't able to get the Details Properly")
		return
	}
	for _, value := range user {
		c.JSON(http.StatusOK, models.Response{
			Message: "Successfully fetched the details",
			Data: models.ApplicantDetail{
				Name:        value.User.Name,
				Email:       value.User.Email,
				PhoneNumber: value.User.PhoneNumber,
				UserId:      value.UserId,
				JobID:       value.JobID,
				Experience:  value.Experience,
				Skills:      value.Skills,
				Language:    value.Language,
				Country:     value.Country,
				JobRole:     value.JobRole,
			},
		})
	}
	loggers.InfoData.Println("Successfully fetched job details")
}

// admin get applied details their jobs by Id
func (service AdminNewHandler) GetPostByJobID(c *gin.Context) {
	parmaJobID := c.Param("job_id")
	jobID, err := strconv.Atoi(parmaJobID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: "Error occurred while converting job_id to integer",
		})
		return
	}

	roleType := c.GetString("role_type")
	roleID := c.GetInt("user_id")

	//Check their roles by admin or users
	if err := helpers.ValidateRoleType(roleType); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Validation failed:", err)
		return
	}

	//get their applied details by their JobIds
	user, err := service.GetAllPostsByJobID(jobID, roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to fetch job details:", err)
		return
	}

	for _, value := range user {
		c.JSON(http.StatusOK, models.Response{
			Message: "Successfully fetched the details",
			Data: models.ApplicantDetail{
				Name:        value.User.Name,
				Email:       value.User.Email,
				PhoneNumber: value.User.PhoneNumber,
				UserId:      value.UserId,
				JobID:       value.JobID,
				Experience:  value.Experience,
				Skills:      value.Skills,
				Language:    value.Language,
				Country:     value.Country,
				JobRole:     value.JobRole,
			},
		})
	}
	loggers.InfoData.Println("Successfully fetched job details")
}

// admin get their jobs by UserId
func (service AdminNewHandler) GetPostByUserID(c *gin.Context) {
	roleType := c.GetString("role_type")
	roleID := c.GetInt("user_id")

	parmID := c.Param("user_id")
	userID, err := strconv.Atoi(parmID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: "invalid user_id parameter",
		})
		return
	}

	//Check their roles by admin or users
	if err := helpers.ValidateRoleType(roleType); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Validation failed:", err)
		return
	}

	//get their User's particular jobs By their userID's
	user, err := service.GetAllPostsByUserID(userID, roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to fetch job applied details:", err)
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "Successfully fetched job details",
		Data:    user,
	})
	loggers.InfoData.Println("Successfully fetched job details")
}

// admin get their own posts by Admin
func (service AdminNewHandler) GetOwnPosts(c *gin.Context) {
	roleType := c.GetString("role_type")
	roleID := c.GetInt("user_id")

	// Check their roles by admin or users
	err := helpers.ValidateRoleType(roleType)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//get thier own post details By admin
	user, err := service.GetOwnPost(roleID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values by admin")
		return
	}

	for _, data := range user {
		c.JSON(http.StatusOK, models.Response{
			Message: "Sucessfully Get the details",
			Data:    data,
		})
	}
	loggers.InfoData.Println("Sucessfully Get the Details by Admin")
}
