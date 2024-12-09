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

type UserHan struct {
	service.UserServices
}

// user or admin get all job details
func (database UserHan) GetAllJobPosts(c *gin.Context) {
	var user []models.JobCreation
	userType := c.GetString("role_type")

	// get their all post details by admin or users
	err := database.GetAllPostsByAdminOrUsers(&user, userType)
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

// user or admin get all jobrole and country
func (database UserHan) GetJobByRole(c *gin.Context) {
	var user []models.JobCreation
	jobs := c.Param("job_title")
	country := c.Param("country")

	userType := c.GetString("role_type")

	//get thier job details by their JobRole
	err := database.GetPostDetailsByTheirRoles(&user, jobs, country, userType)
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

// user or admin get companyName
func (database UserHan) GetByCompanyname(c *gin.Context) {
	var user []models.JobCreation

	CompanyName := c.Param("company_name")
	userType := c.GetString("role_type")

	//users get by thier Company Names by particular Details
	err := database.GetPostDetailsByCompanyNames(&user, CompanyName, userType)
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

// user apply the job in that posts
func (database UserHan) UsersApplyForJobs(c *gin.Context) {
	var user models.UserJobDetails
	var newpost models.JobCreation

	value := c.Param("user_id")
	paramid, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to Invalid request payload")
		return
	}

	tokentype := c.GetString("role_type")
	tokenid := c.GetInt("user_id")

	// Valid their User JobPost with Fields
	err = validation.ValidationUserJob(user, tokentype, tokenid, paramid)
	if err != nil {
		c.JSON(500, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}

	// Check if user ID is applied for the Job or Not
	err = database.CheckJobId(&user, &newpost)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		return
	}

	// user apply the job in that posts
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

// user get by their userowndetails
func (database UserHan) UsersGetTheirDetailsByTheirownIds(c *gin.Context) {
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

	//get their Details by userIds
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
