package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/common/helpers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type UserHan struct {
	service.UserServices
}

// func (database UserHan) MultipleUsers(c *gin.Context) {
// 	var user []models.UserDetails

// 	if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
// 		c.JSON(http.StatusBadRequest, models.CommonResponse{
// 			Error: err.Error()})
// 		loggers.ErrorData.Println("Error occured while getting values")
// 		return
// 	}

// 	datass := database.ServiceFindAllUsers(&user)
// 	if datass != nil {
// 		c.JSON(http.StatusBadRequest, models.CommonResponse{
// 			Error: datass.Error})
// 		loggers.ErrorData.Println("Error occured while getting values")
// 		return
// 	}

// 	for _, vals := range user {
// 		c.JSON(http.StatusOK, models.CommonResponse{
// 			Message: "Sucesssfully get the details",
// 			Data:    vals,
// 		})
// 	}
// }

func (database UserHan) GetAllJobPosts(c *gin.Context) {
	var user []models.JobCreation

	dataOfJobs := database.ServiceGetAllPostDetails(&user)
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

func (database UserHan) GetJobByRole(c *gin.Context) {
	var user []models.JobCreation
	jobs := c.Param("job_title")

	fmt.Println("jobs", jobs)
	dbValues := database.ServiceGetJobDetailsByRole(&user, jobs)
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

func (database UserHan) GetJobByCountry(c *gin.Context) {
	var user []models.JobCreation
	jobs := c.Param("country")

	dbValues := database.ServiceGetDetailsByCountry(&user, jobs)
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

func (database UserHan) ApplyJob(c *gin.Context) {
	var user models.UserJobDetails
	// if err := helpers.CheckuserType(c, "USER"); err != nil {
	// 	c.JSON(http.StatusBadRequest, models.CommonResponse{
	// 		Error: err.Error()})
	// 	loggers.ErrorData.Println("Failed to apply post")
	// 	return
	// }

	value := c.Param("user_id")
	applyuserid, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}

	fmt.Println("values", applyuserid)
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to Invalid request payload")
		return
	}
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	fmt.Println("userid", userid)
	fmt.Println("usertype", userType)
	Dbvalues := database.ApplyJobPost(&user, userType, userid, applyuserid)
	if Dbvalues != nil {
		c.JSON(500, models.CommonResponse{
			Error: "Error occured while creating values"})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}
	c.JSON(http.StatusOK, models.CommonResponse{
		Message: "Sucessfully Applied Job ",
		Data:    user})
}

func (database UserHan) HandlerGetJobAppliedDetailsByUserId(c *gin.Context) {
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
	// userType := c.GetString("role_type")
	// fmt.Println("usertype", userType)

	dbvalues := database.GetJobAppliedDetailsByUserId(&user, values)
	if dbvalues != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: dbvalues.Error})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	for _, valuess := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details",
			Data:    valuess,
		})
	}
}
