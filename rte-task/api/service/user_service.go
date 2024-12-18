package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IUserService interface {
	GetAllJobPosts(company string) ([]models.JobCreation, *models.ErrorResponse)
	GetJobByRole(jobrole string, country string) ([]models.JobCreation, *models.ErrorResponse)
	CreateApplication(userJobDetails *models.UserJobDetails) *models.ErrorResponse
	GetUserAppliedJobs(roleID int) ([]models.UserJobDetails, *models.ErrorResponse)
	GetUserID(userJobDetails *models.UserJobDetails) *models.ErrorResponse
}

type UserService struct {
	repo repository.IUserRepo
}

func InitUserService(db repository.IUserRepo) IUserService {
	return UserService{
		db,
	}
}

// get their all post details by admin or users
func (service UserService) GetAllJobPosts(company string) ([]models.JobCreation, *models.ErrorResponse) {
	return service.repo.GetAllJobPosts(company)
}

// get thier job details by their JobRole by users or admin
func (service UserService) GetJobByRole(jobrole string, country string) ([]models.JobCreation, *models.ErrorResponse) {
	return service.repo.GetJobByRole(jobrole, country)
}

// users apply for the job post
func (service UserService) CreateApplication(userJobDetails *models.UserJobDetails) *models.ErrorResponse {
	return service.repo.CreateApplication(userJobDetails)
}

// users get thier applied details by their own Ids
func (service UserService) GetUserAppliedJobs(roleID int) ([]models.UserJobDetails, *models.ErrorResponse) {
	return service.repo.GetUserAppliedJobs(roleID)
}

// Check if user ID is applied for the Job or Not
func (service UserService) GetUserID(userJobDetails *models.UserJobDetails) *models.ErrorResponse {
	return service.repo.GetUserID(userJobDetails)
}
