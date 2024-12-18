package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IAuthService interface {
	GetSignUpMail(userEmail string) *models.ErrorResponse
	GetSignupNumber(userPhoneNumber string) *models.ErrorResponse
	CreateUser(userDetails *models.UserDetails) *models.ErrorResponse
	GetUserDetail(userDetails *models.UserDetails) (*models.UserDetails, *models.ErrorResponse)
}

type AuthService struct {
	repo repository.IAuthRepo
}

func InitAuthService(db repository.IAuthRepo) IAuthService {
	return AuthService{
		db,
	}
}

// check email is exixts or not in DB
func (service AuthService) GetSignUpMail(userEmail string) *models.ErrorResponse {
	return service.repo.GetSignUpMail(userEmail)
}

// check phone number is exists or not in DB
func (service AuthService) GetSignupNumber(userPhoneNumber string) *models.ErrorResponse {
	return service.repo.GetSignupNumber(userPhoneNumber)
}

// create user details By their roles
func (service AuthService) CreateUser(userDetails *models.UserDetails) *models.ErrorResponse {
	return service.repo.CreateUser(userDetails)
}

// Check Email address while Login with their email ID
func (service AuthService) GetUserDetail(userDetails *models.UserDetails) (*models.UserDetails, *models.ErrorResponse) {
	return service.repo.GetUserDetail(userDetails)
}
