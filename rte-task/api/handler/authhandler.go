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

type AuthConnect struct {
	service.AuthService
}

// create their details
func (database AuthConnect) SignUp(c *gin.Context) {
	var user models.UserDetails
	var count int64

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}
	//signup their each fields
	err := validation.ValidationSignUp(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	//check email is exixts or not in DB
	counts, err := database.CheckEmailIsExists(&user, count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error occured in this email"})
		return
	}
	if counts > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "Email is already registered by someones's ,Try another mail"})
		loggers.ErrorData.Println("This email is already registred and Enter another mail")
		return
	}

	//Hashing the password here
	password := validation.HashPassword(user.Password)
	user.Password = password

	//check phone number is exists or not in DB
	countPhone, err := database.CheckPhoneNumberIsExists(&user, count)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error occured in this PhoneNumber"})
		loggers.ErrorData.Println("Error occured in this PhoneNumber")
		return
	}
	if countPhone > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "PhoneNumber is already registred ,Enter another PhoneNumber"})
		loggers.ErrorData.Println("PhoneNumber is already registred ,Enter another PhoneNumber")
		return
	}

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	//create user details By their roles
	err = database.CreateUserDetails(&user)
	if err != nil {
		c.JSON(500, gin.H{
			"Error": "Can't able to create your data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sucessfully Created the Details",
		"Data":    user})
	loggers.InfoData.Println("Sucessfully Created the Details")
}

// Login with their Details
func (database AuthConnect) Login(c *gin.Context) {
	var user models.UserDetails
	var founduser models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error occured"})
		return
	}

	//Check Email address while Login with their email ID
	err := database.CheckEmailWhileLogin(&user, &founduser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Cant't Find Your MailId"})
		loggers.ErrorData.Println("Cant't Find Your MailId")
		return
	}

	//verify their password is match with signup password
	Password, data := validation.VerifyPassword(user.Password, founduser.Password)
	if !Password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": data})
		return
	}
	//Generate new token here
	token, err := validation.GenerateToken(founduser.Email, founduser.Name, founduser.RoleType, founduser.UserId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Cant't Match Your MailId"})
		loggers.ErrorData.Println("Cant't Match Your MailId")
		return
	}
	user.Token = token

	c.JSON(http.StatusAccepted, gin.H{
		"Message":  "Login Sucessfully",
		"Token":    user.Token,
		"CommonID": founduser.UserId,
		"RoleType": founduser.RoleType,
	})
	loggers.InfoData.Println("Login Sucessfully")
}
