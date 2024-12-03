package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

var validate = validator.New()

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
	if validatonErr := validate.Struct(user); validatonErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": validatonErr.Error()})
		return
	}

	_, err := validation.ValidationFields(user.Email)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error occured in this validation of email"})
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
			"Error": "This email is already registred and Enter another mail"})
		return
	}

	if user.RoleType != "USER" && user.RoleType != "ADMIN" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Role should be Either USER or ADMIN"})
		return
	}

	password := validation.HashPassword(user.Password)
	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Password length should be more than 8"})
		return
	}
	user.Password = password

	if len(user.PhoneNumber) > 10 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": " Phonenumber should be equal to 10"})
		return
	}
	if len(user.PhoneNumber) < 10 {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Phonenumber should be  equal to ten numbers"})
		return
	}

	// countPhone, DbPhone := database.DB.ServicePhoneForm(&user, count)
	// if DbPhone != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"Error": "Error occured in this PhoneNumber"})
	// 	return
	// }
	// if countPhone > 0 {
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"Error": "This phoneNumber is already registred and Enter another PhoneNumber"})
	// 	return
	// }

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	// token, _ := helpers.GenerateToken(user.Email, user.Name, user.RoleType, user.UserId)
	// user.Token = token
	user.RoleId = uuid.New().String()

	Dbvalues := database.ServiceCreate(&user)
	if Dbvalues != nil {
		c.JSON(500, gin.H{
			"Error": "Error occured while creating values"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sucessfully Created the Details",
		"Data":    user})
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
			"Email is incorrect": value.Error})
		return
	}

	Password, data := validation.VerifyPassword(user.Password, founduser.Password)
	if !Password {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": data})
		return
	}
	fmt.Println("founduserid", founduser.UserId)
	token, err := validation.GenerateToken(founduser.Email, founduser.Name, founduser.RoleType, founduser.UserId)
	user.Token = token

	values := database.ServiceFindRoleID(&user, &founduser)

	if values != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message":  "Login Sucessfully",
		"Token":    user.Token,
		"RoleId":   founduser.UserId,
		"RoleType": founduser.RoleType,
	})
}
