package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type UserServices interface {
	ServiceGetAllPostDetails(user *[]models.JobCreation, usertype string) error
	ServiceGetJobDetailsByRole(user *[]models.JobCreation, jobrole string, country string, usertype string) error
	ApplyJobPost(user *models.UserJobDetails, jobtype string, userid int, applyuserid int) error
	GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int, userid int) error
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

func (service *userservice) ApplyJobPost(user *models.UserJobDetails, jobtype string, userid int, applyuserid int) error {
	return service.User.RepoApplyJobPost(user, jobtype, userid, applyuserid)
}

func (service *userservice) GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int, userid int) error {
	return service.User.UserGetJobAppliedDetailsByUserId(user, roleid, userid)
}
