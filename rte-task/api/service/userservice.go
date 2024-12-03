package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type UserServices interface {
	ServiceGetAllPostDetails(user *[]models.JobCreation) error
	ServiceGetJobDetailsByRole(user *[]models.JobCreation, jobrole string) error
	ServiceGetDetailsByCountry(user *[]models.JobCreation, country string) error
	ApplyJobPost(user *models.UserJobDetails, jobtype string, userid int, applyuserid int) error
	GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int) error
}

type userservice struct {
	*repository.UserRepository
}

func (service *userservice) ServiceGetAllPostDetails(user *[]models.JobCreation) error {
	return service.User.RepoGetAllPosts(user)
}

func (service *userservice) ServiceGetJobDetailsByRole(user *[]models.JobCreation, jobrole string) error {
	return service.User.RepoGetByJobRole(user, jobrole)
}

func (service *userservice) ServiceGetDetailsByCountry(user *[]models.JobCreation, country string) error {
	return service.User.RepoGetByCountryDetails(user, country)
}

func (service *userservice) ApplyJobPost(user *models.UserJobDetails, jobtype string, userid int, applyuserid int) error {
	return service.User.RepoApplyJobPost(user, jobtype, userid, applyuserid)
}

func (service *userservice) GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int) error {
	return service.User.UserGetJobAppliedDetailsByUserId(user, roleid)
}
