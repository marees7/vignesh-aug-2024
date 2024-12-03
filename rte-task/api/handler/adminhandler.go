package handler

import (
	"fmt"
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

func (database AdminHand) CreateJobPost(c *gin.Context) {
	var userPost models.JobCreation

	value := c.Param("user_id")
	applyuserid, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	fmt.Println("values", applyuserid)

	if err := c.BindJSON(&userPost); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	_, err = validation.ValidationFields(userPost.CompanyEmail)
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
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	fmt.Println("userid", userid)
	fmt.Println("usertype", userType)

	Dbvalues := database.ServiceCreatePost(&userPost, userType, userid, applyuserid)
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

func (database AdminHand) UpdatePost(c *gin.Context) {
	var post models.JobCreation

	value := c.Param("job_id")
	jobID, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured Invalid job_id_new parameter")
		return
	}

	userids := c.Param("user_id")
	useridvalues, err := strconv.Atoi(userids)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	fmt.Println("useridvalues", useridvalues)
	fmt.Println("jobid", jobID)
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured while request payload")
		return
	}

	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	fmt.Println("userid", userid)
	fmt.Println("usertype", userType)
	if err := database.UpdatePosts(&post, jobID, userType, userid, useridvalues); err != nil {
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

func (database AdminHand) GetJobAppliedDetailsbyrole(c *gin.Context) {
	var user []models.UserJobDetails

	// if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
	// 	c.JSON(http.StatusBadRequest, models.CommonResponse{
	// 		Error: err.Error()})
	// 	loggers.ErrorData.Println("Error occured while getting values")
	// 	return
	// }

	value := c.Param("job_role")

	valueuserid := c.Param("user_id")
	applyuserid, err := strconv.Atoi(valueuserid)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	fmt.Println("values", applyuserid)
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	fmt.Println("userid", userid)
	fmt.Println("usertype", userType)
	dbvalues := database.ServiceGetJobAppliedDetailsbyrole(&user, value, userType, userid, applyuserid)
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

// func (database AdminHand) GetAllAppliedJobDetails(c *gin.Context) {
// 	var user []models.UserJobDetails

// 	// if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
// 	// 	c.JSON(http.StatusBadRequest, models.CommonResponse{
// 	// 		Error: err.Error()})
// 	// 	loggers.ErrorData.Println("Error occured while getting values")
// 	// 	return
// 	// }

// 	value := c.Param("user_id")
// 	applyuserid, err := strconv.Atoi(value)
// 	if err != nil {
// 		c.JSON(404, models.CommonResponse{
// 			Error: err,
// 		})
// 	}
// 	fmt.Println("values", applyuserid)
// 	userType := c.GetString("role_type")
// 	userid := c.GetInt("user_id")

// 	fmt.Println("userid", userid)
// 	fmt.Println("usertype", userType)
// 	dbvalues := database.ServiceGetJobAppliedAllJobs(&user, userType, userid, applyuserid)
// 	if dbvalues != nil {
// 		c.JSON(http.StatusBadRequest, models.CommonResponse{
// 			Error: dbvalues.Error})
// 		loggers.ErrorData.Println("Error occured while getting values")
// 		return
// 	}

// 	for _, values := range user {
// 		c.JSON(http.StatusOK, models.CommonResponse{
// 			Message: "Sucessfully Get the details",
// 			Data:    values,
// 		})
// 	}
// }

func (database AdminHand) GetJobAppliedDetailsByJobId(c *gin.Context) {
	var user []models.UserJobDetails

	// if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
	// 	c.JSON(http.StatusBadRequest, models.CommonResponse{
	// 		Error: err.Error()})
	// 	loggers.ErrorData.Println("Error occured while getting values")
	// 	return
	// }

	value := c.Param("job_id")
	values, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}

	valueuserid := c.Param("user_id")
	applyuserid, err := strconv.Atoi(valueuserid)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	fmt.Println("values", applyuserid)
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	fmt.Println("userid", userid)
	fmt.Println("usertype", userType)
	dbvalues := database.ServiceGetJobAppliedDetailsByJobId(&user, userType, userid, values, applyuserid)
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

func (database AdminHand) GetJobAppliedDetailsByUserId(c *gin.Context) {
	var user []models.UserJobDetails

	// if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
	// 	c.JSON(http.StatusBadRequest, models.CommonResponse{
	// 		Error: err.Error()})
	// 	loggers.ErrorData.Println("Error occured while getting values")
	// 	return
	// }

	value := c.Param("user_id")
	values, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	adminid := c.Param("admin_id")
	adminvalues, err := strconv.Atoi(adminid)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")
	fmt.Println("values", values)
	fmt.Println("useridfromtoken", userid)
	fmt.Println("usertype", userType)

	dbvalues := database.ServiceGetJobAppliedDetailsByUserId(&user, values, userType, userid, adminvalues)
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

// func (database AdminHand) DeletePost(c *gin.Context) {
// 	var post models.JobCreation
// 	value := c.Param("job_id")
// 	values, err := strconv.Atoi(value)
// 	if err != nil {
// 		c.JSON(404, models.CommonResponse{
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
// 	if err := database.DeletedPostsByadmin(&post, values, userType); err != nil {
// 		c.JSON(http.StatusBadRequest, models.CommonResponse{
// 			Error: err.Error(),
// 		})
// 		loggers.ErrorData.Println("Error occured while Deleting that value")
// 		return
// 	}
// 	c.JSON(http.StatusOK, models.CommonResponse{
// 		Message: "Post Deleted successfully",
// 	})
// }
