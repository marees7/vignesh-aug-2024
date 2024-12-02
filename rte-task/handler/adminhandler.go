package handler

import (
	"net/http"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/helpers"
	"github.com/Vigneshwartt/golang-rte-task/loggers"
	"github.com/Vigneshwartt/golang-rte-task/models"
	"github.com/Vigneshwartt/golang-rte-task/validation"
	"github.com/gin-gonic/gin"
)

func (database UserHandler) CreateJobPost(c *gin.Context) {
	var userPost models.JobCreation

	if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")

		return
	}

	if err := c.BindJSON(&userPost); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	_, err := validation.ValidationFields(userPost.CompanyEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: "Error occured in this validation of email"})
		loggers.WarnData.Println("Error occured while getting values")
		return
	}

	if userPost.JobTime != "PART TIME" && userPost.JobTime != "FULL TIME" {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: "Time should be either PART TIME or FULL TIME"})
		loggers.WarnData.Println("Error occured while getting values")
		return
	}

	Dbvalues := database.DB.ServiceCreatePost(&userPost)
	if Dbvalues != nil {
		c.JSON(500, models.CommonResponse{
			Error: "Error occured while creating values"})
		loggers.ErrorData.Println("Error occured while creating values")
		return
	}
	c.JSON(http.StatusOK, models.CommonResponse{
		Message: "Sucessfully Created the Details",
		Data:    userPost})
}

func (database UserHandler) UpdatePost(c *gin.Context) {
	var post models.JobCreation

	value := c.Param("job_id_new")
	jobID, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured Invalid job_id_new parameter")
		return
	}

	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured while request payload")
		return
	}

	userType := c.GetString("role_type")

	if err := database.DB.UpdatePosts(&post, jobID, userType); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to update post")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Message":   "Post updated successfully",
		"JobStatus": post.JobStatus,
		"JobTime":   post.JobTime,
		"Vacancy":   post.Vacancy,
	})
}

func (database UserHandler) GetJobAppliedDetailsbyrole(c *gin.Context) {
	var user []models.UserJobDetails

	if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	value := c.Param("job_role")
	dbvalues := database.DB.ServiceGetJobAppliedDetailsbyrole(&user, value)
	if dbvalues != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: dbvalues.Error})
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

func (database UserHandler) GetAllAppliedJobDetails(c *gin.Context) {
	var user []models.UserJobDetails

	if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	dbvalues := database.DB.ServiceGetJobAppliedAllJobs(&user)
	if dbvalues != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: dbvalues.Error})
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

func (database UserHandler) GetJobAppliedDetailsByJobId(c *gin.Context) {
	var user []models.UserJobDetails

	if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	value := c.Param("job_id")
	values, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	dbvalues := database.DB.ServiceGetJobAppliedDetailsByJobId(&user, values)
	if dbvalues != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: dbvalues.Error})
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

func (database UserHandler) GetJobAppliedDetailsByUserId(c *gin.Context) {
	var user []models.UserJobDetails

	if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
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
	dbvalues := database.DB.ServiceGetJobAppliedDetailsByUserId(&user, values)
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

// func (database UserHandler) DeletePost(c *gin.Context) {
// 	var post models.JobCreation
// 	value := c.Param("job_id_new")
// 	values, err := strconv.Atoi(value)
// 	if err != nil {
// 		c.JSON(404,models.CommonResponse{
// 			Error: err.Error(),
// 		})
// 		loggers.ErrorData.Println("Error occured in this form")
// 		return
// 	}
// 	if err := c.Bind(&post); err != nil {
// 		c.JSON(http.StatusBadRequest, models.CommonResponse{
// 			Error: err.Error(),
// 		})
// 		loggers.ErrorData.Println("Error occured while binding values")
// 		return
// 	}
// 	userType := c.GetString("role_type")
// 	if err := database.DB.DeletedPostsByadmin(&post, values, userType); err != nil {
// 		c.JSON(http.StatusBadRequest,models.CommonResponse{
// 			Error: err.Error(),
// 		})
// 		loggers.ErrorData.Println("Error occured while Deleting that value")
// 		return
// 	}
// 	c.JSON(http.StatusOK,models.CommonResponse{
// 		Message: "Post Deleted successfully",
// 	})
// }
