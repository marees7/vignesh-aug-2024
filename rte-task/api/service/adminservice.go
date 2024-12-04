package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type AdminService interface {
	ServiceFindAllUsers(user *[]models.UserDetails) error
	ServiceCreatePost(user *models.JobCreation, userType string, userid int, applyuserid int) error
	UpdatePosts(user *models.JobCreation, jobid int, roletype string, userid int, useridvalues int) error
	ServiceGetJobAppliedDetailsbyrole(user *[]models.UserJobDetails, roletype string, usertype string, userid int, applyuserid int) error
	ServiceGetJobAppliedDetailsByJobId(user *[]models.UserJobDetails, usertype string, userid int, roleid int, applyuserid int) error
	ServiceGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int, usertype string, userid int, adminid int) error
	ServiceGetPostedDetailsByAdmin(user *[]models.JobCreation, usertype string, userid int, adminid int) error
}
type adminservice struct {
	*repository.UserRepository
}

func (service *adminservice) ServiceFindAllUsers(user *[]models.UserDetails) error {
	return service.Admin.RepoFindAllUsers(user)
}

func (service *adminservice) ServiceCreatePost(user *models.JobCreation, userType string, userid int, applyuserid int) error {
	return service.Admin.RepoCreateNewPost(user, userType, userid, applyuserid)
}

func (service *adminservice) UpdatePosts(user *models.JobCreation, jobid int, roletype string, userid int, useridvalues int) error {
	return service.Admin.RepoUpdateJobPost(user, jobid, roletype, userid, useridvalues)
}
func (service *adminservice) ServiceGetJobAppliedDetailsbyrole(user *[]models.UserJobDetails, roletype string, usertype string, userid int, applyuserid int) error {
	return service.Admin.RepoGetJobAppliedDetailsbyrole(user, roletype, usertype, userid, applyuserid)
}

func (service *adminservice) ServiceGetJobAppliedDetailsByJobId(user *[]models.UserJobDetails, usertype string, userid int, roleid int, applyuserid int) error {
	return service.Admin.RepoGetJobAppliedDetailsByJobId(user, usertype, userid, roleid, applyuserid)
}
func (service *adminservice) ServiceGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int, usertype string, userid int, adminid int) error {
	return service.Admin.RepoGetJobAppliedDetailsByUserId(user, roleid, usertype, userid, adminid)
}
func (service *adminservice) ServiceGetPostedDetailsByAdmin(user *[]models.JobCreation, usertype string, userid int, adminid int) error {
	return service.Admin.RepoGetPostedDetailsByAdmin(user, usertype, userid, adminid)
}
