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

func (database AuthConnect) SignUp(c *gin.Context) {
	var user models.UserDetails
	var count int64

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}
	err := validation.ValidationSignUp(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	counts, err := database.ServiceRepoemail(&user, count)
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

	password := validation.HashPassword(user.Password)
	user.Password = password

	countPhone, DbPhone := database.ServicePhoneForm(&user, count)
	if DbPhone != nil {
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

	Dbvalues := database.ServiceCreate(&user)
	if Dbvalues != nil {
		c.JSON(500, gin.H{
			"Error": "Can't able to create your data"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sucessfully Created the Details",
		"Data":    user})
	loggers.InfoData.Println("Sucessfully Created the Details")
}

func (database AuthConnect) Login(c *gin.Context) {
	var user models.UserDetails
	var founduser models.UserDetails

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error occured"})
		return
	}

	value := database.ServiceLoginEmail(&user, &founduser)
	if value != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Cant't Find Your MailId"})
		loggers.ErrorData.Println("Cant't Find Your MailId")
		return
	}

	Password, data := validation.VerifyPassword(user.Password, founduser.Password)
	if !Password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": data})
		return
	}
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
