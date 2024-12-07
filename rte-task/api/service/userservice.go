package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type UserServices interface {
	ServiceGetAllPostDetails(user *[]models.JobCreation, usertype string) error
	ServiceGetJobDetailsByRole(user *[]models.JobCreation, jobrole string, country string, usertype string) error
	ApplyJobPost(user *models.UserJobDetails) error
	GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int, userid int) error
	CheckJobId(user *models.UserJobDetails, newpost *models.JobCreation) error
	ServiceGetByCompanyName(user *[]models.JobCreation, company string, usertype string) error
}

type userservice struct {
	*repository.UserRepository
}

func (service *userservice) ServiceGetAllPostDetails(user *[]models.JobCreation, usertype string) error {
	return service.User.RepoGetAllPosts(user, usertype)
}

func (service *userservice) ServiceGetJobDetailsByRole(user *[]models.JobCreation, jobrole string, country string, usertype string) error {
	return service.User.RepoGetByJobRole(user, jobrole, country, usertype)
}
func (service *userservice) ServiceGetByCompanyName(user *[]models.JobCreation, company string, usertype string) error {
	return service.User.RepoGetByCompanyName(user, company, usertype)
}

func (service *userservice) ApplyJobPost(user *models.UserJobDetails) error {
	return service.User.RepoApplyJobPost(user)
}

func (service *userservice) GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int, userid int) error {
	return service.User.UserGetJobAppliedDetailsByUserId(user, roleid, userid)
}
func (service *userservice) CheckJobId(user *models.UserJobDetails, newpost *models.JobCreation) error {
	return service.User.CheckUserJobId(user, newpost)
}
