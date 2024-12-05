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

	adminId := c.Param("admin_id")
	paramid, err := strconv.Atoi(adminId)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: err,
		})
	}
	fmt.Println("payloadid", paramid)

	if err := c.BindJSON(&userPost); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Cant able to get the JobPost Details")
		return
	}
	fmt.Println("checkdetails", userPost)
	tokentype := c.GetString("role_type")
	tokenid := c.GetInt("user_id")

	fmt.Println("userid", tokenid)
	fmt.Println("usertype", tokentype)

	err = validation.ValidationJobPost(userPost, paramid, tokenid, tokentype)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}
	Dbvalues := database.ServiceCreatePost(&userPost)
	if Dbvalues != nil {
		c.JSON(500, models.CommonResponse{
			Error: Dbvalues.Error()})
		loggers.ErrorData.Println("OOPS! Your Id or roletype is Mismatching Here,Check It Once Again")
		return
	}

	c.JSON(http.StatusOK, models.CommonResponse{
		Message: "Sucessfully Created the Details",
		Data:    userPost})
	loggers.InfoData.Println("Sucessfully Created the Details")
}

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
	useridvalues, err := strconv.Atoi(adminid)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: "Error occured while String Convertion,Please check properly",
		})
	}
	// fmt.Println("useridvalues", useridvalues)
	// fmt.Println("jobid", jobID)
	if err := c.ShouldBindJSON(&post); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Error occured while request payload")
		return
	}

	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	// fmt.Println("userid", userid)
	// fmt.Println("usertype", userType)
	err = validation.ValidationAdminFields(post, userType, userid, useridvalues)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	if err := database.UpdatePosts(&post, jobID); err != nil {
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
	loggers.InfoData.Println("Sucessfully Updated the Details")

}

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
	fmt.Println("values", useridvalues)
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	fmt.Println("roletype", jobrole)
	err = validation.ValidationCheck(userType, userid, useridvalues)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}
	err = database.ServiceGetJobAppliedDetailsbyrole(&user, jobrole)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Can't able to get the Details Properly")
		return
	}

	for _, values := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details",
			Data:    values,
		})
	}
	loggers.InfoData.Println("Sucessfully Get the Details")
}

func (database AdminHand) GetJobAppliedDetailsByJobId(c *gin.Context) {
	var user []models.UserJobDetails

	value := c.Param("job_id")
	jobId, err := strconv.Atoi(value)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: "Error occured while String Convertion,Please check properly",
		})
	}

	valueuserid := c.Param("admin_id")
	applyuserid, err := strconv.Atoi(valueuserid)
	if err != nil {
		c.JSON(404, models.CommonResponse{
			Error: "Error occured while String Convertion,Please check properly",
		})
	}
	fmt.Println("values", applyuserid)
	userType := c.GetString("role_type")
	userid := c.GetInt("user_id")

	fmt.Println("userid", userid)
	fmt.Println("usertype", userType)
	err = validation.ValidationCheck(userType, userid, applyuserid)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}
	err = database.ServiceGetJobAppliedDetailsByJobId(&user, jobId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured while getting values")
		return
	}

	for _, values := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details",
			Data:    values,
		})
	}
	loggers.InfoData.Println("Sucessfully Get the Details")
}

func (database AdminHand) GetJobAppliedDetailsByUserId(c *gin.Context) {
	var user []models.UserJobDetails

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
	err = validation.ValidationCheck(userType, userid, adminvalues)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}
	err = database.ServiceGetJobAppliedDetailsByUserId(&user, values)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Cant able to get the job Applied Details By theirs ID")
		return
	}

	for _, valuess := range user {
		c.JSON(http.StatusOK, models.CommonResponse{
			Message: "Sucessfully Get the details",
			Data:    valuess,
		})
	}
	loggers.InfoData.Println("Sucessfully Get the Details by their IDs")
}

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
	fmt.Println("adminid", adminId)
	fmt.Println("useridfromtoken", userid)
	fmt.Println("usertype", userType)

	err = validation.ValidationCheck(userType, userid, adminId)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this validation of postdetails")
		return
	}

	err = database.ServiceGetPostedDetailsByAdmin(&user, adminId)
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
