package repository

import (
	"fmt"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IAuthRepo interface {
	CreateUser(userDetails *models.UserDetails) *dto.ErrorResponse
	GetUserMail(userEmail string) *dto.ErrorResponse
	GetUserPhoneNumber(userPhoneNumber string) *dto.ErrorResponse
	GetUserDetail(userDetails *models.UserDetails) (*models.UserDetails, *dto.ErrorResponse)
}

type AuthRepo struct {
	*internals.NewConnection
}

func InitAuthRepo(db *internals.NewConnection) IAuthRepo {
	return &AuthRepo{
		db,
	}
}

// create user details By their roles
func (database *AuthRepo) CreateUser(userDetails *models.UserDetails) *dto.ErrorResponse {
	db := database.Create(userDetails)
	if db.Error != nil {
		return &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Errorf("can't able to create your data"),
		}
	}

	return nil
}

// check email is exixts or not in DB
func (database *AuthRepo) GetUserMail(userEmail string) *dto.ErrorResponse {
	var count int64

	db := database.Model(&models.UserDetails{}).
		Where("email=?", userEmail).
		Count(&count)
	if db.Error != nil {
		return &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("error occured while searching the email "),
		}
	}

	if count > 0 {
		return &dto.ErrorResponse{
			StatusCode: http.StatusAlreadyReported,
			Error:      fmt.Errorf(" error Email is already registered by someones's ,Try another mail"),
		}
	}

	return nil
}

// check phone number is exists or not in DB
func (database *AuthRepo) GetUserPhoneNumber(userPhoneNumber string) *dto.ErrorResponse {
	var count int64

	db := database.Model(&models.UserDetails{}).
		Where("phone_number=?", userPhoneNumber).
		Count(&count)
	if db.Error != nil {
		return &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("cant' able to create your Phone number"),
		}
	}

	if count > 0 {
		return &dto.ErrorResponse{
			StatusCode: http.StatusAlreadyReported,
			Error:      fmt.Errorf(" error Phone Number is already registered by someones's ,Please verify it"),
		}
	}

	return nil
}

// Check Email address while Login with their email ID
func (database *AuthRepo) GetUserDetail(userDetails *models.UserDetails) (*models.UserDetails, *dto.ErrorResponse) {
	var founduser *models.UserDetails

	db := database.
		Where("email=?", userDetails.Email).
		First(&founduser)
	if db.Error != nil {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("cant' Match your MailId properly,check it once"),
		}
	}

	return founduser, nil
}
