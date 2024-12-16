package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type I_AuthService interface {
	GetLoginEmail(user *models.UserDetails) error
	CreateUser(user *models.UserDetails) error
	GetUserMail(user *models.UserDetails, founduser *models.UserDetails) error
	GetSignupNumber(user *models.UserDetails) error
}

type AuthService struct {
	repository.I_AuthRepo
}

func GetAuthService(db repository.I_AuthRepo) I_AuthService {
	return &AuthService{
		db,
	}
}

// check email is exixts or not in DB
func (repo *AuthService) GetLoginEmail(user *models.UserDetails) error {
	return repo.GetValidEmailAddress(user)
}

// check phone number is exists or not in DB
func (repo *AuthService) GetSignupNumber(user *models.UserDetails) error {
	return repo.GetPhoneNumber(user)
}

// create user details By their roles
func (repo *AuthService) CreateUser(user *models.UserDetails) error {
	return repo.CreateUsers(user)
}

// Check Email address while Login with their email ID
func (repo *AuthService) GetUserMail(user *models.UserDetails, founduser *models.UserDetails) error {
	return repo.GetUserMails(user, founduser)
}
