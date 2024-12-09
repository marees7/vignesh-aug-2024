package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type AuthService interface {
	CheckEmailIsExists(user *models.UserDetails, count int64) (int64, error)
	CreateUserDetails(user *models.UserDetails) error
	CheckEmailWhileLogin(user *models.UserDetails, founduser *models.UserDetails) error
	CheckPhoneNumberIsExists(user *models.UserDetails, count int64) (int64, error)
}

type authservice struct {
	*repository.UserRepository
}

// check email is exixts or not in DB
func (service *authservice) CheckEmailIsExists(user *models.UserDetails, count int64) (int64, error) {
	return service.Auth.CheckEmailAddress(user, count)
}

// check phone number is exists or not in DB
func (service *authservice) CheckPhoneNumberIsExists(user *models.UserDetails, count int64) (int64, error) {
	return service.Auth.CheckPhoneNumber(user, count)
}

// create user details By their roles
func (service *authservice) CreateUserDetails(user *models.UserDetails) error {
	return service.Auth.CreateDetailsByTheirRoles(user)
}

// Check Email address while Login with their email ID
func (service *authservice) CheckEmailWhileLogin(user *models.UserDetails, founduser *models.UserDetails) error {
	return service.Auth.LoginEmailCheckExists(user, founduser)
}
