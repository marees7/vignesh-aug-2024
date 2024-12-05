package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type AdminService interface {
	ServiceFindAllUsers(user *[]models.UserDetails) error
	ServiceCreatePost(user *models.JobCreation) error
	UpdatePosts(user *models.JobCreation, jobid int) error
	ServiceGetJobAppliedDetailsbyrole(user *[]models.UserJobDetails, roletype string) error
	ServiceGetJobAppliedDetailsByJobId(user *[]models.UserJobDetails, jobid int) error
	ServiceGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int) error
	ServiceGetPostedDetailsByAdmin(user *[]models.JobCreation, adminid int) error
}
type adminservice struct {
	*repository.UserRepository
}

func (service *adminservice) ServiceFindAllUsers(user *[]models.UserDetails) error {
	return service.Admin.RepoFindAllUsers(user)
}

func (service *adminservice) ServiceCreatePost(user *models.JobCreation) error {
	return service.Admin.RepoCreateNewPost(user)
}

func (service *adminservice) UpdatePosts(user *models.JobCreation, jobid int) error {
	return service.Admin.RepoUpdateJobPost(user, jobid)
}
func (service *adminservice) ServiceGetJobAppliedDetailsbyrole(user *[]models.UserJobDetails, jobrole string) error {
	return service.Admin.RepoGetJobAppliedDetailsbyrole(user, jobrole)
}

func (service *adminservice) ServiceGetJobAppliedDetailsByJobId(user *[]models.UserJobDetails, jobid int) error {
	return service.Admin.RepoGetJobAppliedDetailsByJobId(user, jobid)
}
func (service *adminservice) ServiceGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int) error {
	return service.Admin.RepoGetJobAppliedDetailsByUserId(user, roleid)
}
func (service *adminservice) ServiceGetPostedDetailsByAdmin(user *[]models.JobCreation, adminid int) error {
	return service.Admin.RepoGetPostedDetailsByAdmin(user, adminid)
}
