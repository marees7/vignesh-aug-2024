package handler

import (
	"net/http"
	"strconv"

	"github.com/Vigneshwartt/golang-rte-task/helpers"
	"github.com/Vigneshwartt/golang-rte-task/loggers"
	"github.com/Vigneshwartt/golang-rte-task/models"
	"github.com/gin-gonic/gin"
)

func (database UserHandler) ApplyJob(c *gin.Context) {
	var user models.UserJobDetails

	if err := helpers.CheckuserType(c, "USER"); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error()})
		loggers.ErrorData.Println("Failed to apply post")
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.CommonResponse{
			Error: err.Error(),
		})
		loggers.ErrorData.Println("Failed to Invalid request payload")
		return
	}
	userType := c.GetString("role_type")

	Dbvalues := database.DB.ApplyJobPost(&user, userType)
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

func (database UserHandler) HandlerGetJobAppliedDetailsByUserId(c *gin.Context) {
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
	dbvalues := database.DB.GetJobAppliedDetailsByUserId(&user, values)
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
