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

var validate = validator.New()

type UserHandler struct {
	DB service.UserService
}

func NewHandlerRepository(db service.UserService) UserHandler {
	return UserHandler{DB: db}
}

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

	// _, err = validation.ValidRoleType(user.RoleType)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"Error": "Role should be Either USER or ADMIN"})
	// 	return
	// }
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

	// _, err = validation.ValidatePhoneNumber(user.PhoneNumber)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"Error": "Error occured in this validation of Phonenumber"})
	// 	return
	// }

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
	countPhone, DbPhone := database.DB.ServicePhoneForm(&user, count)
	if DbPhone != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error occured in this PhoneNumber"})
		return
	}
	if countPhone > 0 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": "This phoneNumber is already registred and Enter another PhoneNumber"})
		return
	}

	user.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	user.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	token, _ := helpers.GenerateToken(user.Email, user.Name, user.RoleType, user.RoleId)
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

	_, err := helpers.GenerateToken(founduser.Email, founduser.Name, founduser.RoleType, founduser.RoleId)
	values := database.DB.ServiceFindRoleID(&user, &founduser)

	if values != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, gin.H{
		"Message":  "Login Sucessfully",
		"Token":    founduser.Token,
		"RoleId":   founduser.RoleId,
		"RoleType": founduser.RoleType,
	})
}

func (database UserHandler) GetUserss(c *gin.Context) {
	var user []models.UsersTable

	if err := helpers.CheckuserType(c, "ADMIN"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": err.Error()})
		return
	}

	datass := database.DB.ServiceFindAllUsers(&user)
	if datass != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Errors": datass.Error})
		return
	}

	for _, vals := range user {
		c.JSON(http.StatusOK, gin.H{
			"Message": "Sucesssfully get the details",
			"Data":    vals,
		})
	}
}

func (database UserHandler) GetUser(c *gin.Context) {
	var user models.UsersTable
	roleid := c.Param("role_id")

	if err := helpers.MatchUserType(c, roleid); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	dbValue := database.DB.ServiceFindSpecificUser(&user, roleid)
	if dbValue != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Errors": dbValue.Error})
		return
	}
	c.JSON(http.StatusAccepted, user)
}

func (database UserHandler) CreateJobPost(c *gin.Context) {
	var userPost models.JobCreation

	if err := c.BindJSON(&userPost); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error()})
		return
	}

	_, err := validation.ValidationFields(userPost.CompanyEmail)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error occured in this validation of email"})
		return
	}

	if userPost.JobTime != "PART TIME" && userPost.JobTime != "FULL TIME" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Time should be either PART TIME or FULL TIME"})
		return
	}

	Dbvalues := database.DB.ServiceCreatePost(&userPost)
	if Dbvalues != nil {
		c.JSON(500, gin.H{
			"Error": "Error occured while creating values"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Sucessfully Created the Details",
		"Data":    userPost})
}
