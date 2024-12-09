package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type AdminService interface {
	CreatePostForUsers(user *models.JobCreation) error
	UpdatePosts(user *models.JobCreation, jobid int, adminId int) error
	GetJobAppliedDetailsbyRole(user *[]models.UserJobDetails, roletype string, adminid int) error
	GetAppliedDetailsByJobId(user *[]models.UserJobDetails, jobid int, adminid int) error
	GetPostDetailsByUserId(user *[]models.UserJobDetails, roleid int, adminvalues int) error
	GetPostDetailsByAdmin(user *[]models.JobCreation, adminid int) error
}
type adminservice struct {
	*repository.UserRepository
}

// create their Jobposts
func (service *adminservice) CreatePostForUsers(user *models.JobCreation) error {
	return service.Admin.CreatePostDetailsByAdmin(user)
}

// update their job post by their IDs
func (service *adminservice) UpdatePosts(user *models.JobCreation, jobid int, adminID int) error {
	return service.Admin.UpdateJobPostsByAdmin(user, jobid, adminID)
}

// get their postdetails jobDetailsBy role
func (service *adminservice) GetJobAppliedDetailsbyRole(user *[]models.UserJobDetails, jobrole string, adminid int) error {
	return service.Admin.GetDetailsByRoleByAdmin(user, jobrole, adminid)
}

// get their applied details by their JobIds
func (service *adminservice) GetAppliedDetailsByJobId(user *[]models.UserJobDetails, jobid int, adminid int) error {
	return service.Admin.GetJobDetailsByJobIdByAdmin(user, jobid, adminid)
}

// get their User's particular jobs By their userID's
func (service *adminservice) GetPostDetailsByUserId(user *[]models.UserJobDetails, roleid int, adminvalues int) error {
	return service.Admin.GetJobDetailsByUserIdByAdmin(user, roleid, adminvalues)
}

// get thier own post details By admin
func (service *adminservice) GetPostDetailsByAdmin(user *[]models.JobCreation, adminid int) error {
	return service.Admin.GetOwnPostDetailsByAdmin(user, adminid)
}
