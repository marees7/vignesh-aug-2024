package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IUserService interface {
	CreateApplication(userJobDetails *models.UserJobDetails) *dto.ErrorResponse
	GetAllJobPosts(searchJobs map[string]interface{}) ([]models.JobCreation, *dto.ErrorResponse, int64)
	GetUserAppliedJobs(userJobs map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64)
	GetUserByID(userJobDetails *models.UserJobDetails) *dto.ErrorResponse
}

type UserService struct {
	repo repository.IUserRepo
}

func InitUserService(db repository.IUserRepo) IUserService {
	return &UserService{
		db,
	}
}

// users apply for the job post
func (service *UserService) CreateApplication(userJobDetails *models.UserJobDetails) *dto.ErrorResponse {
	return service.repo.CreateApplication(userJobDetails)
}

// get their all post details by admin or users
func (service *UserService) GetAllJobPosts(searchJobs map[string]interface{}) ([]models.JobCreation, *dto.ErrorResponse, int64) {
	return service.repo.GetAllJobPosts(searchJobs)
}

// users get thier applied details by their own Ids
func (service *UserService) GetUserAppliedJobs(userJobs map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64) {
	return service.repo.GetUserAppliedJobs(userJobs)
}

// Check if user ID is applied for the Job or Not
func (service *UserService) GetUserByID(userJobDetails *models.UserJobDetails) *dto.ErrorResponse {
	return service.repo.GetUserByID(userJobDetails)
}
