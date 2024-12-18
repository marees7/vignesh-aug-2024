package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IAdminService interface {
	CreateJobPost(jobCreation *models.JobCreation) *models.ErrorResponse
	UpdateJobPost(jobCreation *models.JobCreation, jobid int) *models.ErrorResponse
	GetPostByRole(jobRole string, userID int) ([]models.ApplicantDetail, *models.ErrorResponse)
	GetPostByJobID(jobID int, userID int) ([]models.ApplicantDetail, *models.ErrorResponse)
	GetPostByUserID(applicantID int, userID int) ([]models.UserJobDetails, *models.ErrorResponse)
	GetPosts(userID int) ([]models.JobCreation, *models.ErrorResponse)
}
type AdminService struct {
	repo repository.IAdminRepo
}

func InitAdminService(db repository.IAdminRepo) IAdminService {
	return AdminService{
		db,
	}
}

// create their Jobposts
func (service AdminService) CreateJobPost(jobCreation *models.JobCreation) *models.ErrorResponse {
	return service.repo.CreateJobPost(jobCreation)
}

// update their job post by their IDs
func (service AdminService) UpdateJobPost(jobCreation *models.JobCreation, jobid int) *models.ErrorResponse {
	return service.repo.UpdateJobPost(jobCreation, jobid)
}

// get their postdetails jobDetailsBy role
func (service AdminService) GetPostByRole(jobRole string, userID int) ([]models.ApplicantDetail, *models.ErrorResponse) {

	userJobdetails, err := service.repo.GetPostByRole(jobRole, userID)
	if err != nil {
		return nil, err
	}
	applicantDetails := make([]models.ApplicantDetail, 0)
	for _, value := range userJobdetails {
		applicantDetails = append(applicantDetails, models.ApplicantDetail{
			Name:        value.User.Name,
			Email:       value.User.Email,
			PhoneNumber: value.User.PhoneNumber,
			UserId:      value.UserId,
			JobID:       value.JobID,
			Experience:  value.Experience,
			Skills:      value.Skills,
			Language:    value.Language,
			Country:     value.Country,
			JobRole:     value.JobRole,
		})
	}
	return applicantDetails, nil
}

func (service AdminService) GetPostByJobID(jobID int, userID int) ([]models.ApplicantDetail, *models.ErrorResponse) {
	userJobDetails, err := service.repo.GetPostByJobID(jobID, userID)
	if err != nil {
		return nil, err
	}

	applicantDetails := make([]models.ApplicantDetail, 0)
	for _, value := range userJobDetails {
		applicantDetails = append(applicantDetails, models.ApplicantDetail{
			Name:        value.User.Name,
			Email:       value.User.Email,
			PhoneNumber: value.User.PhoneNumber,
			UserId:      value.UserId,
			JobID:       value.JobID,
			Experience:  value.Experience,
			Skills:      value.Skills,
			Language:    value.Language,
			Country:     value.Country,
			JobRole:     value.JobRole,
		})
	}
	return applicantDetails, nil
}

// get their User's particular jobs By their userID's
func (service AdminService) GetPostByUserID(applicantID int, userID int) ([]models.UserJobDetails, *models.ErrorResponse) {
	return service.repo.GetPostByUserID(applicantID, userID)
}

// get thier own post details By admin
func (service AdminService) GetPosts(userID int) ([]models.JobCreation, *models.ErrorResponse) {
	return service.repo.GetPosts(userID)
}
