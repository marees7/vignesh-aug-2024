package repository

import (
	"fmt"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IAuthRepo interface {
	GetSignUpMail(userEmail string) *models.ErrorResponse
	GetSignupNumber(userPhoneNumber string) *models.ErrorResponse
	CreateUser(userDetails *models.UserDetails) *models.ErrorResponse
	GetUserDetail(userDetails *models.UserDetails) (*models.UserDetails, *models.ErrorResponse)
}

type AuthRepo struct {
	*internals.ConnectionNew
}

func InitAuthRepo(db *internals.ConnectionNew) IAuthRepo {
	return AuthRepo{
		db,
	}
}

// check email is exixts or not in DB
func (database AuthRepo) GetSignUpMail(userEmail string) *models.ErrorResponse {
	var count int64

	db := database.Model(&models.UserDetails{}).
		Where("email=?", userEmail).
		Count(&count)
	if db.Error != nil {
		return &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("error occured while searching the email "),
		}
	}
	if count > 0 {
		return &models.ErrorResponse{
			StatusCode: http.StatusAlreadyReported,
			Error:      fmt.Errorf(" error Email is already registered by someones's ,Try another mail"),
		}
	}
	return nil
}

// check phone number is exists or not in DB
func (database AuthRepo) GetSignupNumber(userPhoneNumber string) *models.ErrorResponse {
	var count int64

	db := database.Model(&models.UserDetails{}).
		Where("phone_number=?", userPhoneNumber).
		Count(&count)
	if db.Error != nil {
		return &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("cant' able to create your Phone number"),
		}
	}
	if count > 0 {
		return &models.ErrorResponse{
			StatusCode: http.StatusAlreadyReported,
			Error:      fmt.Errorf(" error Phone Number is already registered by someones's ,Please verify it"),
		}
	}
	return nil
}

// create user details By their roles
func (database AuthRepo) CreateUser(userDetails *models.UserDetails) *models.ErrorResponse {
	db := database.Create(userDetails)
	if db.Error != nil {
		return &models.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Errorf("can't able to create your data"),
		}
	}
	return nil
}

// Check Email address while Login with their email ID
func (database AuthRepo) GetUserDetail(userDetails *models.UserDetails) (*models.UserDetails, *models.ErrorResponse) {
	var founduser *models.UserDetails
	db := database.
		Where("email=?", userDetails.Email).
		First(&founduser)
	if db.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("cant' Match your MailId properly,check it once"),
		}
	}
	return founduser, nil
}
