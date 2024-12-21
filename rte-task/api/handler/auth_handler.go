package handler

import (
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/common/helpers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service service.IAuthService
}

// create their details
func (handler AuthHandler) CreateUser(c *gin.Context) {
	var userDetail models.UserDetails

	if err := c.BindJSON(&userDetail); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{
			Error: err.Error()})
		loggers.WarnData.Println("Can't able to Bind the data", err)
		return
	}

	//signup their each fields
	err := validation.ValidateSignUp(userDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("can't able to validate the user ", err)
		return
	}

	//check email is exixts or not in DB
	errorResponse := handler.Service.GetUserMail(userDetail.Email)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error Occured", errorResponse.Error)
		return
	}

	//Hashing the password here
	password := helpers.HashPassword(userDetail.Password)
	userDetail.Password = password

	//check phone number is exists or not in DB
	errorResponse = handler.Service.GetUserPhoneNumber(userDetail.PhoneNumber)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured in this PhoneNumber", errorResponse.Error)
		return
	}

	//create user details By their roles
	errorResponse = handler.Service.CreateUser(&userDetail)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Can't create the user", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Sucessfully Created the User Detail", userDetail.UserID)
	c.JSON(http.StatusCreated, models.Response{
		Message: "Sucessfully Created the User Detail",
		Data:    userDetail})
}

// GetUserDetail with their Details
func (handler AuthHandler) GetUserDetail(c *gin.Context) {
	var userDetail models.UserDetails

	if err := c.BindJSON(&userDetail); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Can't Get the Details", err.Error())
		return
	}

	//Check Email address while Login with their email ID
	userLogin, errorResponse := handler.Service.GetUserDetail(&userDetail)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Cant't Find Your MailId", errorResponse.Error)
		return
	}

	//verify their password is match with signup password
	password, err := validation.VerifyPassword(userDetail.Password, userLogin.Password)
	if !password {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured", err.Error())
		return
	}

	//Generate new token here
	token, err := validation.GenerateToken(userLogin.Email, userLogin.Name, userLogin.RoleType, userLogin.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Cant't able to Generate token ,check it", err)
		return
	}
	// userDetails.Token = token

	loggers.InfoData.Println("Login Sucessfully", userLogin.UserID)
	c.JSON(http.StatusOK, models.LoginUser{
		Message:  "Login Sucessfully",
		Token:    token,
		ID:       userLogin.UserID,
		RoleType: userLogin.RoleType,
	})
}
