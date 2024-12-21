package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IAdminService interface {
	CreateJobPost(jobCreation *models.JobCreation) *dto.ErrorResponse
	GetApplicantAndJobDetails(jobDetailsMap map[string]interface{}) ([]*dto.ApplicantDetail, *dto.ErrorResponse, int64)
	GetJobsAppliedByUser(userIDJobs map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64)
	GetJobsCreated(createdPosts map[string]interface{}) ([]models.JobCreation, *dto.ErrorResponse, int64)
	UpdateJobPost(jobData *models.JobCreation, jobID int, userID int) *dto.ErrorResponse
	DeleteJobPost() *dto.ErrorResponse
}
type AdminService struct {
	repo repository.IAdminRepo
}

func InitAdminService(db repository.IAdminRepo) IAdminService {
	return &AdminService{
		db,
	}
}

// create their Jobposts
func (service *AdminService) CreateJobPost(jobCreation *models.JobCreation) *dto.ErrorResponse {
	return service.repo.CreateJobPost(jobCreation)
}

// get their postdetails jobDetailsBy role
func (service *AdminService) GetApplicantAndJobDetails(jobDetailsMap map[string]interface{}) ([]*dto.ApplicantDetail, *dto.ErrorResponse, int64) {
	userJobdetails, err, count := service.repo.GetApplicantAndJobDetails(jobDetailsMap)
	if err != nil {
		return nil, err, 0
	}

	applicantDetails := make([]*dto.ApplicantDetail, 0)
	for _, value := range userJobdetails {
		applicantDetails = append(applicantDetails, &dto.ApplicantDetail{
			Name:        value.User.Name,
			Email:       value.User.Email,
			PhoneNumber: value.User.PhoneNumber,
			UserID:      value.UserID,
			JobID:       *value.JobID,
			Experience:  value.Experience,
			Skills:      value.Skills,
			Language:    value.Language,
			Country:     value.Country,
			JobRole:     value.JobRole,
		})
	}

	return applicantDetails, nil, count
}

// get their User's particular jobs By their userID's
func (service *AdminService) GetJobsAppliedByUser(userIDJobs map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64) {
	return service.repo.GetJobsAppliedByUser(userIDJobs)
}

// get thier own post details By admin
func (service *AdminService) GetJobsCreated(createdPosts map[string]interface{}) ([]models.JobCreation, *dto.ErrorResponse, int64) {
	return service.repo.GetJobsCreated(createdPosts)
}

// update their job post by their IDs
func (service *AdminService) UpdateJobPost(jobData *models.JobCreation, jobID int, userID int) *dto.ErrorResponse {
	return service.repo.UpdateJobPost(jobData, jobID, userID)
}

func (service *AdminService) DeleteJobPost() *dto.ErrorResponse {
	return service.repo.DeleteJobPost()
}
