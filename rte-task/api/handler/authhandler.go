package handler

import (
	"net/http"
	"time"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	service.I_AuthService
}

// create their details
func (service AuthHandler) SignUp(c *gin.Context) {
	var user models.UserDetails
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		return
	}
	//signup their each fields
	err := validation.ValidationSignUp(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: err.Error()})
		return
	}

	//check email is exixts or not in DB
	err = service.GetLoginEmail(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		return
	}

	//Hashing the password here
	password := validation.HashPassword(user.Password)
	user.Password = password

	//check phone number is exists or not in DB
	err = service.GetSignupNumber(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured in this PhoneNumber")
		return
	}

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	//create user details By their roles
	err = service.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		return
	}

	c.JSON(http.StatusOK, models.Response{
		Message: "Sucessfully Created the Details",
		Data:    user})
	loggers.InfoData.Println("Sucessfully Created the Details")
}

// Login with their Details
func (service AuthHandler) Login(c *gin.Context) {
	var user models.UserDetails
	var founduser models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: "Error occured while Binding the data"})
		return
	}

	//Check Email address while Login with their email ID
	err := service.GetUserMail(&user, &founduser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Cant't Find Your MailId")
		return
	}

	//verify their password is match with signup password
	Password, data := validation.VerifyPassword(user.Password, founduser.Password)
	if !Password {
		c.JSON(http.StatusBadRequest, models.Response{
			Error: data})
		return
	}
	//Generate new token here
	token, err := validation.GenerateToken(founduser.Email, founduser.Name, founduser.RoleType, founduser.UserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			Error: "Cant't able to Generate token ,check it"})
		loggers.ErrorData.Println("Cant't able to Generate token ,check it")
		return
	}
	user.Token = token

	c.JSON(http.StatusOK, gin.H{
		"Message":  "Login Sucessfully",
		"Token":    user.Token,
		"ID":       founduser.UserId,
		"RoleType": founduser.RoleType,
	})
	loggers.InfoData.Println("Login Sucessfully")
}
