package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IAuthService interface {
	CreateUser(userDetails *models.UserDetails) *models.ErrorResponse
	GetUserMail(userEmail string) *models.ErrorResponse
	GetUserPhoneNumber(userPhoneNumber string) *models.ErrorResponse
	GetUserDetail(userDetails *models.UserDetails) (*models.UserDetails, *models.ErrorResponse)
}

type AuthService struct {
	repo repository.IAuthRepo
}

func InitAuthService(db repository.IAuthRepo) IAuthService {
	return &AuthService{
		db,
	}
}

// create user details By their roles
func (service *AuthService) CreateUser(userDetails *models.UserDetails) *models.ErrorResponse {
	return service.repo.CreateUser(userDetails)
}

// check email is exixts or not in DB
func (service *AuthService) GetUserMail(userEmail string) *models.ErrorResponse {
	return service.repo.GetUserMail(userEmail)
}

// check phone number is exists or not in DB
func (service *AuthService) GetUserPhoneNumber(userPhoneNumber string) *models.ErrorResponse {
	return service.repo.GetUserPhoneNumber(userPhoneNumber)
}

// Check Email address while Login with their email ID
func (service *AuthService) GetUserDetail(userDetails *models.UserDetails) (*models.UserDetails, *models.ErrorResponse) {
	return service.repo.GetUserDetail(userDetails)
}
