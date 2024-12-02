package handler

import (
	"net/http"
	"time"

	"github.com/Vigneshwartt/golang-rte-task/helpers"
	"github.com/Vigneshwartt/golang-rte-task/models"
	"github.com/Vigneshwartt/golang-rte-task/service"
	"github.com/Vigneshwartt/golang-rte-task/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

type UserHandler struct {
	DB service.UserService
}

func NewHandlerRepository(db service.UserService) UserHandler {
	return UserHandler{DB: db}
}

var validate = validator.New()

func (database UserHandler) SignUp(c *gin.Context) {
	var user models.UsersTable
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

	counts, err := database.DB.ServiceRepoemail(&user, count)
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
	token, _ := helpers.GenerateToken(user.Email, user.Name, user.RoleType, user.UserId)
	user.Token = token
	user.RoleId = uuid.New().String()

	Dbvalues := database.DB.ServiceCreate(&user)
	if Dbvalues != nil {
		c.JSON(500, gin.H{
			"Error": "Error occured while creating values"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sucessfully Created the Details",
		"Data":    user})
}

func (database UserHandler) Login(c *gin.Context) {
	var user models.UsersTable
	var founduser models.UsersTable

	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error occured"})
		return
	}

	value := database.DB.ServiceLoginEmail(&user, &founduser)
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

	_, err := helpers.GenerateToken(founduser.Email, founduser.Name, founduser.RoleType, founduser.UserId)
	values := database.DB.ServiceFindRoleID(&user, &founduser)

	if values != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message":  "Login Sucessfully",
		"Token":    founduser.Token,
		"RoleId":   founduser.UserId,
		"RoleType": founduser.RoleType,
	})
}
