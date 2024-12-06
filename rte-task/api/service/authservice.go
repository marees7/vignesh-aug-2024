package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type AuthService interface {
	ServiceRepoemail(user *models.UserDetails, count int64) (int64, error)
	ServiceCreate(user *models.UserDetails) error
	ServiceLoginEmail(user *models.UserDetails, founduser *models.UserDetails) error
	ServicePhoneForm(user *models.UserDetails, count int64) (int64, error)
}

type authservice struct {
	*repository.UserRepository
}

func (service *authservice) ServiceRepoemail(user *models.UserDetails, count int64) (int64, error) {
	return service.Auth.RepoEmailForm(user, count)
}

func (service *authservice) ServicePhoneForm(user *models.UserDetails, count int64) (int64, error) {
	return service.Auth.RepoPhoneForm(user, count)
}

func (service *authservice) ServiceCreate(user *models.UserDetails) error {
	return service.Auth.RepoCreate(user)
}

func (service *authservice) ServiceLoginEmail(user *models.UserDetails, founduser *models.UserDetails) error {
	return service.Auth.RepoLoginEmail(user, founduser)
}
