package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type I_UserService interface {
	GetAllPosts() ([]models.JobCreation, error)
	GetPostByRoles(jobrole string, country string) ([]models.JobCreation, error)
	CreateApplicationJob(user *models.UserJobDetails) error
	GetUserJobs(roleID int) ([]models.UserJobDetails, error)
	GetUserID(user *models.UserJobDetails) error
	GetByCompany(company string) ([]models.JobCreation, error)
}

type Userservice struct {
	repository.I_UserRepo
}

func GetUserService(db repository.I_UserRepo) I_UserService {
	return &Userservice{
		db,
	}
}

// get their all post details by admin or users
func (repo *Userservice) GetAllPosts() ([]models.JobCreation, error) {
	return repo.GetAllPost()
}

// get thier job details by their JobRole by users or admin
func (repo *Userservice) GetPostByRoles(jobrole string, country string) ([]models.JobCreation, error) {
	return repo.GetPostsByRoles(jobrole, country)
}

// get by thier Company Names by particular Details by users or admin
func (repo *Userservice) GetByCompany(company string) ([]models.JobCreation, error) {
	return repo.GetByCompanys(company)
}

// users apply for the job post
func (repo *Userservice) CreateApplicationJob(user *models.UserJobDetails) error {
	return repo.CreateJobApplication(user)
}

// users get thier applied details by their own Ids
func (repo *Userservice) GetUserJobs(roleID int) ([]models.UserJobDetails, error) {
	return repo.GetUserJob(roleID)
}

// Check if user ID is applied for the Job or Not
func (repo *Userservice) GetUserID(user *models.UserJobDetails) error {
	return repo.GetUserId(user)
}
