package handler

import (
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/api/service"
	"github.com/Vigneshwartt/golang-rte-task/api/validation"
	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/common/helpers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/loggers"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	Service service.IAuthService
}

// @Summary signup
// @Description Sign up their details
// @Param user body models.UserDetails true "User"
// @Tags Auth
// @Accept json
// @Produce json
// @Success 201 {object}  models.UserDetails
// @Failure 500 {object} dto.Response
// @Failure 422 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 208 {object} dto.Response
// @Router /v1/auth/signup [post]
func (handler AuthHandler) CreateUser(c *gin.Context) {
	var userDetail models.UserDetails

	if err := c.BindJSON(&userDetail); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.Response{
			Error: err.Error()})
		loggers.WarnData.Println("Can't able to Bind the data-", err)
		return
	}

	//signup their each fields
	err := validation.ValidateSignUp(userDetail)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("can't able to validate the user-", err)
		return
	}

	//check email is exixts or not in DB
	errorResponse := handler.Service.GetUserMail(userDetail.Email)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, dto.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error Occured when fetching user mail-", errorResponse.Error)
		return
	}

	//Hashing the password here
	password := helpers.HashPassword(userDetail.Password)
	userDetail.Password = password

	//check phone number is exists or not in DB
	errorResponse = handler.Service.GetUserPhoneNumber(userDetail.PhoneNumber)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, dto.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Error occured when fetching user PhoneNumber-", errorResponse.Error)
		return
	}

	//create user details By their roles
	errorResponse = handler.Service.CreateUser(&userDetail)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, dto.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Can't create the user-", errorResponse.Error)
		return
	}

	loggers.InfoData.Println("Sucessfully Created the User Detail-", userDetail.UserID)
	c.JSON(http.StatusCreated, dto.Response{
		Message: "Sucessfully Created the User Detail",
		Data:    userDetail})
}

// @Summary Loginuser
// @Description  Login with their Details
// @Param user body models.UserDetails true "User"
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object}  dto.LoginUser
// @Failure 500 {object} dto.Response
// @Failure 404 {object} dto.Response
// @Failure 400 {object} dto.Response
// @Router /v1/auth/login [post]
func (handler AuthHandler) GetUserDetail(c *gin.Context) {
	var userDetail models.UserDetails

	if err := c.BindJSON(&userDetail); err != nil {
		c.JSON(http.StatusUnprocessableEntity, dto.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Can't Get the Details-", err.Error())
		return
	}

	//Check Email address while Login with their email ID
	userLogin, errorResponse := handler.Service.GetUserDetail(&userDetail)
	if errorResponse != nil {
		c.JSON(errorResponse.StatusCode, dto.Response{
			Error: errorResponse.Error.Error()})
		loggers.ErrorData.Println("Cant't Find Your MailID-", errorResponse.Error)
		return
	}

	//verify their password is match with signup password
	password, err := validation.VerifyPassword(userDetail.Password, userLogin.Password)
	if !password {
		c.JSON(http.StatusBadRequest, dto.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Error occured when verifying password-", err.Error())
		return
	}

	//Generate new token here
	token, err := validation.GenerateToken(userLogin.Email, userLogin.Name, userLogin.RoleType, userLogin.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Error: err.Error()})
		loggers.ErrorData.Println("Cant't able to Generate token ,check it-", err)
		return
	}
	// userDetails.Token = token

	loggers.InfoData.Println("Login Sucessfully", userLogin.UserID)
	c.JSON(http.StatusOK, dto.LoginUser{
		Message:  "Login Sucessfully",
		Token:    token,
		ID:       userLogin.UserID,
		RoleType: userLogin.RoleType,
	})
}
