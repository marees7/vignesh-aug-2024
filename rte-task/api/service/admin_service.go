package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IAdminService interface {
	CreateJobPost(jobCreation *models.JobCreation) *models.ErrorResponse
	GetApplicantAndJobDetails(jobDetailsMap map[string]interface{}) ([]models.ApplicantDetail, *models.ErrorResponse, int64)
	GetJobsAppliedByUser(userIDJobs map[string]interface{}) ([]models.UserJobDetails, *models.ErrorResponse, int64)
	GetJobsCreated(createdPosts map[string]int) ([]models.JobCreation, *models.ErrorResponse, int64)
	UpdateJobPost(jobUpdation *models.JobCreation, jobID int, userID int) *models.ErrorResponse
	DeleteJobPost() *models.ErrorResponse
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
func (service *AdminService) CreateJobPost(jobCreation *models.JobCreation) *models.ErrorResponse {
	return service.repo.CreateJobPost(jobCreation)
}

// get their postdetails jobDetailsBy role
func (service *AdminService) GetApplicantAndJobDetails(jobDetailsMap map[string]interface{}) ([]models.ApplicantDetail, *models.ErrorResponse, int64) {
	userJobdetails, err, count := service.repo.GetApplicantAndJobDetails(jobDetailsMap)
	if err != nil {
		return nil, err, 0
	}

	applicantDetails := make([]models.ApplicantDetail, 0)
	for _, value := range userJobdetails {
		applicantDetails = append(applicantDetails, models.ApplicantDetail{
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
func (service *AdminService) GetJobsAppliedByUser(userIDJobs map[string]interface{}) ([]models.UserJobDetails, *models.ErrorResponse, int64) {
	return service.repo.GetJobsAppliedByUser(userIDJobs)
}

// get thier own post details By admin
func (service *AdminService) GetJobsCreated(createdPosts map[string]int) ([]models.JobCreation, *models.ErrorResponse, int64) {
	return service.repo.GetJobsCreated(createdPosts)
}

// update their job post by their IDs
func (service *AdminService) UpdateJobPost(jobUpdation *models.JobCreation, jobID int, userID int) *models.ErrorResponse {
	return service.repo.UpdateJobPost(jobUpdation, jobID, userID)
}

func (service *AdminService) DeleteJobPost() *models.ErrorResponse {
	return service.repo.DeleteJobPost()
}
