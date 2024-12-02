package handler

import (
	"fmt"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/helpers"
	"github.com/Vigneshwartt/golang-rte-task/loggers"
	"github.com/Vigneshwartt/golang-rte-task/models"
	"github.com/gin-gonic/gin"
)

func (database UserHandler) MultipleUsers(c *gin.Context) {
	var user []models.UsersTable

	if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	datass := database.DB.ServiceFindAllUsers(&user)
	if datass != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: datass.Error})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	for _, vals := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucesssfully get the details",
			Data:    vals,
		})
	}
}

func (database UserHandler) GetUser(c *gin.Context) {
	var user models.UsersTable
	roleid := c.Param("role_id")

	if err := helpers.MatchUserType(c, roleid); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	dbValue := database.DB.ServiceFindSpecificUser(&user, roleid)
	if dbValue != nil {
		c.JSON(http.StatusInternalServerError, models.CommonResponse{
			Error: dbValue.Error})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}
	c.JSON(http.StatusAccepted, user)
}

func (database UserHandler) GetAllJobPosts(c *gin.Context) {
	var user []models.JobCreation

	dataOfJobs := database.DB.ServiceGetAllPostDetails(&user)
	if dataOfJobs != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: dataOfJobs.Error})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}
	for _, values := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details",
			Data:    values,
		})
	}
}

func (database UserHandler) GetJobByRole(c *gin.Context) {
	var user []models.JobCreation
	jobs := c.Param("job_title")

	fmt.Println("jobs", jobs)
	dbValues := database.DB.ServiceGetJobDetailsByRole(&user, jobs)
	if dbValues != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: dbValues.Error})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}
	for _, values := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details",
			Data:    values,
		})
	}
}

func (database UserHandler) GetJobByCountry(c *gin.Context) {
	var user []models.JobCreation
	jobs := c.Param("country")

	dbValues := database.DB.ServiceGetDetailsByCountry(&user, jobs)
	if dbValues != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: dbValues.Error})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}
	for _, values := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details",
			Data:    values,
		})
	}
}
