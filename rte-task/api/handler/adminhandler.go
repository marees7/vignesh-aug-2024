package handler

import (
	"net/http"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type AdminHand struct {
	service.AdminService
}

// admin creates new JobPost
func (database AdminHand) CreateJobPost(c *gin.Context) {
	var userPost models.JobCreation

	adminId := c.Param("admin_id")
	paramid, err := strconv.Atoi(adminId)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}

	if err := c.BindJSON(&userPost); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Cant able to get the JobPost Details")
		return
	}
	tokentype := c.GetString("role_type")
	tokenid := c.GetInt("user_id")

	//validate each fields in Jobpost
	err = validation.ValidationJobPost(userPost, paramid, tokenid, tokentype)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//create their Jobposts
	err = database.CreatePostForUsers(&userPost)
	if err != nil {
		c.JSON(500, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("OOPS! Your Id or roletype is Mismatching Here,Check It Once Again")
		return
	}

	c.JSON(http.StatusOK, models.CommonResponse{
		Message: "Sucessfully Created the Details",
		Data:    userPost})
	loggers.InfoData.Println("Sucessfully Created the Details")
}

// admin updates their own JobPost
func (database AdminHand) UpdatePost(c *gin.Context) {
	var post models.JobCreation

	value := c.Param("job_id")
	jobID, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: "Error occured while String Convertion,Please check properly",
		})
		loggers.ErrorData.Println("Error occured Invalid JobId Found,Please Check")
		return
	}

	adminid := c.Param("admin_id")
	adminIdvalues, err := strconv.Atoi(adminid)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: "Error occured while String Convertion,Please check properly",
		})
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured while request payload")
		return
	}

	//Validate each fields
	err = validation.ValidationUpdatePost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	// Valid their Admin JobPosts with Fields
	err = validation.ValidationAdminFields(post, userType, userid, adminIdvalues)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//update their job post by their IDs
	if err := database.UpdatePosts(&post, jobID, adminIdvalues); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
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
func (database AdminHand) GetJobAppliedDetailsbyrole(c *gin.Context) {
	var user []models.UserJobDetails

	jobrole := c.Param("job_role")
	valueuserid := c.Param("admin_id")
	useridvalues, err := strconv.Atoi(valueuserid)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	//Valid their roles by admin or users
	err = validation.ValidationCheck(userType, userid, useridvalues)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//get their jobDetailsBy role
	err = database.GetJobAppliedDetailsbyRole(&user, jobrole, useridvalues)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Can't able to get the Details Properly")
		return
	}
	var response []gin.H
	for _, details := range user {
		userDetails := gin.H{}
		if details.User != nil {
			userDetails = gin.H{
				"name":         details.User.Name,
				"email":        details.User.Email,
				"phone_number": details.User.PhoneNumber,
			}
		}

		response = append(response, gin.H{
			"user_id":    details.UserId,
			"job_id":     details.JobID,
			"experience": details.Experience,
			"skills":     details.Skills,
			"language":   details.Language,
			"country":    details.Country,
			"job_role":   details.JobRole,
			"created_at": details.CreatedAt,
			"updated_at": details.UpdatedAt,
			"user":       userDetails,
		})
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Successfully fetched the details",
			Data:    response,
		})
		loggers.InfoData.Println("Successfully fetched job details")
	}
}

// admin get applied details their jobs by Id
func (database AdminHand) GetJobAppliedDetailsByJobId(c *gin.Context) {
	var user []models.UserJobDetails

	value := c.Param("job_id")
	jobId, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: "Error occurred while converting job_id to integer",
		})
		return
	}

	valueUserId := c.Param("admin_id")
	applyUserId, err := strconv.Atoi(valueUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: "Error occurred while converting admin_id to integer",
		})
		return
	}

	userType := c.GetString("role_type")
	userId := c.GetInt("user_id")

	//Valid their roles by admin or users
	if err := validation.ValidationCheck(userType, userId, applyUserId); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Validation failed:", err)
		return
	}

	//get their applied details by their JobIds
	if err := database.GetAppliedDetailsByJobId(&user, jobId, applyUserId); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to fetch job details:", err)
		return
	}

	var response []gin.H
	for _, details := range user {
		userDetails := gin.H{}
		if details.User != nil {
			userDetails = gin.H{
				"name":         details.User.Name,
				"email":        details.User.Email,
				"phone_number": details.User.PhoneNumber,
			}
		}

		response = append(response, gin.H{
			"user_id":    details.UserId,
			"job_id":     details.JobID,
			"experience": details.Experience,
			"skills":     details.Skills,
			"language":   details.Language,
			"country":    details.Country,
			"job_role":   details.JobRole,
			"created_at": details.CreatedAt,
			"updated_at": details.UpdatedAt,
			"user":       userDetails,
		})
	}

	c.JSON(http.StatusOK, models.CommonResponse{
		Message: "Successfully fetched the details",
		Data:    response,
	})
	loggers.InfoData.Println("Successfully fetched job details")
}

// admin get their jobs by UserId
func (database AdminHand) GetJobAppliedDetailsByUserId(c *gin.Context) {
	var user []models.UserJobDetails
	adminid := c.Param("admin_id")
	adminvalues, err := strconv.Atoi(adminid)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}

	userType := c.GetString("role_type")
	userIdnew := c.GetInt("user_id")

	userIdParam := c.Param("user_id")
	userId, err := strconv.Atoi(userIdParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: "invalid user_id parameter",
		})
		return
	}

	//Valid their roles by admin or users
	if err := validation.ValidationCheck(userType, userIdnew, adminvalues); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Validation failed:", err)
		return
	}

	//get their User's particular jobs By their userID's
	if err := database.GetPostDetailsByUserId(&user, userId, adminvalues); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to fetch job applied details:", err)
		return
	}

	c.JSON(http.StatusOK, models.CommonResponse{
		Message: "Successfully fetched job details",
		Data:    user,
	})
	loggers.InfoData.Println("Successfully fetched job details")
}

// admin get their own posts by Admin
func (database AdminHand) GetJobsByAdmin(c *gin.Context) {
	var user []models.JobCreation

	adminid := c.Param("admin_id")
	adminId, err := strconv.Atoi(adminid)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	// Valid their roles by admin or users
	err = validation.ValidationCheck(userType, userid, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	//get thier own post details By admin
	err = database.GetPostDetailsByAdmin(&user, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values by admin")
		return
	}

	for _, valuess := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details",
			Data:    valuess,
		})
	}
	loggers.InfoData.Println("Sucessfully Get the Details by Admin")
}
