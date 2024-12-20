package repository

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type IAdminRepo interface {
	CreateJobPost(jobCreation *models.JobCreation) *models.ErrorResponse
	GetApplicantAndJobDetails(jobDetails map[string]interface{}) ([]models.UserJobDetails, *models.ErrorResponse, int64)
	GetJobsAppliedByUser(userIDJobs map[string]interface{}) ([]models.UserJobDetails, *models.ErrorResponse, int64)
	GetJobsCreated(createdPosts map[string]int) ([]models.JobCreation, *models.ErrorResponse, int64)
	UpdateJobPost(jobUpdation *models.JobCreation, jobID int, userID int) *models.ErrorResponse
	DeleteJobPost() *models.ErrorResponse
}
type AdminRepo struct {
	*internals.ConnectionNew
}

func InitAdminRepo(db *internals.ConnectionNew) IAdminRepo {
	return &AdminRepo{
		db,
	}
}

// Admin creates new post in this case
func (database *AdminRepo) CreateJobPost(jobCreation *models.JobCreation) *models.ErrorResponse {
	db := database.Create(jobCreation)

	if errors.Is(db.Error, gorm.ErrInvalidData) {
		return &models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      db.Error,
		}
	}

	return nil
}

func (database *AdminRepo) GetApplicantAndJobDetails(jobDetails map[string]interface{}) ([]models.UserJobDetails, *models.ErrorResponse, int64) {
	var userJobDetails []models.UserJobDetails
	var count int64

	jobRole := jobDetails["JobRole"].(string)
	jobID := jobDetails["JobID"].(int)
	userID := jobDetails["UserID"].(int)
	limit := jobDetails["Limit"].(int)
	offset := jobDetails["Offset"].(int)

	query := database.Debug()

	if jobRole != "" {
		result := database.
			Where(&models.UserJobDetails{JobRole: jobRole}).
			First(&userJobDetails)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil, &models.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the job applications for Job Role"),
				}, 0
			}
			return nil, &models.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      result.Error,
			}, 0
		}

		query = query.Debug().Preload("User").
			Where("job_role=? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ?)", jobRole, userID).
			Find(&userJobDetails)
	}

	if jobID != 0 {
		result := database.
			Where("job_id=?", jobID).
			First(&userJobDetails)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil, &models.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the job applications for Job ID"),
				}, 0
			}
			return nil, &models.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      result.Error,
			}, 0
		}

		query = query.Preload("User").
			Where("job_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ? )", jobID, userID).
			Find(&userJobDetails)
	}

	result := query.Debug().Limit(limit).Offset(offset).Find(&userJobDetails).Count(&count)
	if result.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      result.Error,
		}, 0
	} else if len(userJobDetails) == 0 {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("your not access to view this Details"),
		}, 0
	}

	return userJobDetails, nil, count
}

// Admin get Job applied details by USER ID
func (database *AdminRepo) GetJobsAppliedByUser(userIDJobs map[string]interface{}) ([]models.UserJobDetails, *models.ErrorResponse, int64) {
	var userJobdetails []models.UserJobDetails
	var count int64

	limit := userIDJobs["Limit"].(int)
	offset := userIDJobs["Offset"].(int)
	applicantID := userIDJobs["Applicant"].(int)
	userID := userIDJobs["UserID"].(int)

	data := database.Debug().Preload("Job").
		Where("user_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ?)", applicantID, userID).
		Limit(limit).Offset(offset).
		Find(&userJobdetails).Count(&count)

	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("failed to fetch job details: %s", data.Error),
		}, 0
	} else if len(userJobdetails) == 0 {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("no job details found for the given user"),
		}, 0
	}

	return userJobdetails, nil, count
}

// Admin get their Own Post details
func (database *AdminRepo) GetJobsCreated(createdPosts map[string]int) ([]models.JobCreation, *models.ErrorResponse, int64) {
	var jobCreation []models.JobCreation
	var count int64

	userID := createdPosts["UserID"]
	limit := createdPosts["Limit"]
	offset := createdPosts["Offset"]

	dbvalue := database.
		Where("admin_id = ?", userID).Limit(limit).Offset(offset).
		Find(&jobCreation).Count(&count)

	if errors.Is(dbvalue.Error, gorm.ErrRecordNotFound) {
		return nil, &models.ErrorResponse{
			Error:      dbvalue.Error,
			StatusCode: http.StatusBadRequest,
		}, 0
	}

	return jobCreation, nil, count
}

// Admin updates their job posts
func (database *AdminRepo) UpdateJobPost(jobUpdation *models.JobCreation, jobID int, userID int) *models.ErrorResponse {
	var jobUpdate models.JobCreation

	db := database.
		Where("job_id = ?", jobID).
		First(&jobUpdate)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      db.Error,
		}
	}

	updatDb := database.Model(&models.JobCreation{}).Where("job_id = ?", jobID).Updates(map[string]interface{}{
		"JobStatus": jobUpdation.JobStatus,
		"Vacancy":   jobUpdation.Vacancy,
	})

	if errors.Is(updatDb.Error, gorm.ErrInvalidData) {
		return &models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      db.Error,
		}
	} else if updatDb.RowsAffected == 0 {
		return &models.ErrorResponse{
			StatusCode: http.StatusBadGateway,
			Error:      fmt.Errorf("can't able to update jobPost"),
		}
	}

	return nil
}

func (database *AdminRepo) DeleteJobPost() *models.ErrorResponse {
	var jobDeletion []models.JobCreation

	result := database.Debug().Select("created_at").Where("created_at < ?", "2024-01-01 00:00:00").Find(&jobDeletion).Delete(&jobDeletion)
	if result.Error != nil {
		return &models.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      result.Error,
		}
	} else if result.RowsAffected == 0 {
		return &models.ErrorResponse{
			StatusCode: http.StatusNotModified,
			Error:      fmt.Errorf("no data are deleted here"),
		}
	}

	return nil
}
