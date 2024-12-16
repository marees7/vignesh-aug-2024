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

type Userhandler struct {
	service.I_UserService
}

// user or admin get all job details
func (service Userhandler) GetAllJobPosts(c *gin.Context) {
	var err error
	var user []models.JobCreation

	country := c.Query("company_name")

	// Call the service to fetch data
	if country != "" {
		user, err = service.GetByCompany(country)
	} else {
		user, err = service.GetAllPosts()
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
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
func (service Userhandler) GetJobByRole(c *gin.Context) {
	paramJobTitle := c.Param("job_title")
	paramJobCountry := c.Param("country")

	//get thier job details by their JobRole
	user, err := service.GetPostByRoles(paramJobTitle, paramJobCountry)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}
	for _, value := range user {
		c.JSON(http.StatusOK, models.Response{
			Message: "Sucessfully Get the details by their JobRole",
			Data:    value,
		})
	}
	loggers.InfoData.Println("Sucessfully Get the JobDetails By thier roles")
}

// user apply the job in that posts
func (service Userhandler) CreateApplication(c *gin.Context) {
	var user models.UserJobDetails

	paramUserId := c.Param("user_id")
	UserId, err := strconv.Atoi(paramUserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err,
		})
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to Invalid request payload")
		return
	}

	roleType := c.GetString("role_type")
	roleID := c.GetInt("user_id")

	// Valid their User JobPost with Fields
	err = validation.ValidationUserJob(user, roleType, roleID, UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}

	// Check if user ID is applied for the Job or Not
	err = service.GetUserID(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		return
	}

	// user apply the job in that posts
	err = service.CreateApplicationJob(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully Applied Job ",
		Data:    user})
	loggers.InfoData.Println("Sucessfully Applied the Job")
}

// user get by their userowndetails
func (service Userhandler) GetUsersDetails(c *gin.Context) {
	roleType := c.GetString("role_type")
	roleID := c.GetInt("user_id")

	if err := helpers.ValidateUserType(roleType); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	//get their Details by userIds
	user, err := service.GetUserJobs(roleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	for _, value := range user {
		c.JSON(http.StatusOK, models.Response{
			Message: "Sucessfully Get the details by thier own userIds",
			Data:    value,
		})
	}
	loggers.InfoData.Println("Sucessfully Get the details by thier own userIds")
}
