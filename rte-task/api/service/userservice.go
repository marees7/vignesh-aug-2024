package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type UserServices interface {
	GetAllPostsByAdminOrUsers(user *[]models.JobCreation, usertype string) error
	GetPostDetailsByTheirRoles(user *[]models.JobCreation, jobrole string, country string, usertype string) error
	ApplyJobPost(user *models.UserJobDetails) error
	GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int, userid int) error
	CheckJobId(user *models.UserJobDetails, newpost *models.JobCreation) error
	GetPostDetailsByCompanyNames(user *[]models.JobCreation, company string, usertype string) error
}

type userservice struct {
	*repository.UserRepository
}

// get their all post details by admin or users
func (service *userservice) GetAllPostsByAdminOrUsers(user *[]models.JobCreation, usertype string) error {
	return service.User.GetAllPostsByAllUsers(user, usertype)
}

// get thier job details by their JobRole by users or admin
func (service *userservice) GetPostDetailsByTheirRoles(user *[]models.JobCreation, jobrole string, country string, usertype string) error {
	return service.User.GetAllPostDetailsByTheirRoles(user, jobrole, country, usertype)
}

// get by thier Company Names by particular Details by users or admin
func (service *userservice) GetPostDetailsByCompanyNames(user *[]models.JobCreation, company string, usertype string) error {
	return service.User.GetAllPostDetailsByCompanyNames(user, company, usertype)
}

// users apply for the job post
func (service *userservice) ApplyJobPost(user *models.UserJobDetails) error {
	return service.User.UsersApplyForTheJobPosts(user)
}

// users get thier applied details by their own Ids
func (service *userservice) GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, roleid int, userid int) error {
	return service.User.UserGetJobAppliedDetailsByUserId(user, roleid, userid)
}

// Check if user ID is applied for the Job or Not
func (service *userservice) CheckJobId(user *models.UserJobDetails, newpost *models.JobCreation) error {
	return service.User.CheckUserJobId(user, newpost)
}
