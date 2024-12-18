package handler

import (
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service service.IAuthService
}

// create their details
func (handler AuthHandler) CreateUser(c *gin.Context) {
	var userDetails models.UserDetails

	if err := c.BindJSON(&userDetails); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{
			Error: err.Error()})
		return
	}

	//signup their each fields
	err := validation.ValidationSignUp(userDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		return
	}

	//check email is exixts or not in DB
	errorResponse := handler.Service.GetSignUpMail(userDetails.Email)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		return
	}

	//Hashing the password here
	password := validation.HashPassword(userDetails.Password)
	userDetails.Password = password

	//check phone number is exists or not in DB
	errorResponse = handler.Service.GetSignupNumber(userDetails.PhoneNumber)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured in this PhoneNumber")
		return
	}

	//create user details By their roles
	errorResponse = handler.Service.CreateUser(&userDetails)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		return
	}

	c.JSON(http.StatusCreated, models.Response{
		Message: "Sucessfully Created the Details",
		Data:    userDetails})
	loggers.InfoData.Println("Sucessfully Created the Details")
}

// GetUserDetail with their Details
func (handler AuthHandler) GetUserDetail(c *gin.Context) {
	var userDetails models.UserDetails

	if err := c.BindJSON(&userDetails); err != nil {
		c.JSON(http.StatusUnprocessableEntity, models.Response{
			Error: "Error occured while Binding the data"})
		return
	}

	//Check Email address while Login with their email ID
	foundUser, errorResponse := handler.Service.GetUserDetail(&userDetails)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, models.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Cant't Find Your MailId")
		return
	}

	//verify their password is match with signup password
	password, data := validation.VerifyPassword(userDetails.Password, foundUser.Password)
	if !password {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: data})
		return
	}

	//Generate new token here
	token, err := validation.GenerateToken(foundUser.Email, foundUser.Name, foundUser.RoleType, foundUser.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Cant't able to Generate token ,check it"})
		loggers.ErrorData.Println("Cant't able to Generate token ,check it")
		return
	}
	userDetails.Token = token

	c.JSON(http.StatusOK, gin.H{
		"Message":  "Login Sucessfully",
		"Token":    userDetails.Token,
		"ID":       foundUser.UserId,
		"RoleType": foundUser.RoleType,
	})
	loggers.InfoData.Println("Login Sucessfully")
}
