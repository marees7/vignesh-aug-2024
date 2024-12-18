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
	UpdateJobPost(jobCreation *models.JobCreation, jobID int) *models.ErrorResponse
	GetPostByRole(roleType string, userID int) ([]models.UserJobDetails, *models.ErrorResponse)
	GetPostByJobID(jobID int, userID int) ([]models.UserJobDetails, *models.ErrorResponse)
	GetPostByUserID(applicantID int, userID int) ([]models.UserJobDetails, *models.ErrorResponse)
	GetPosts(userID int) ([]models.JobCreation, *models.ErrorResponse)
}
type AdminRepo struct {
	*internals.ConnectionNew
}

func InitAdminRepo(db *internals.ConnectionNew) IAdminRepo {
	return AdminRepo{
		db,
	}
}

// Admin creates new post in this case
func (database AdminRepo) CreateJobPost(jobCreation *models.JobCreation) *models.ErrorResponse {
	db := database.Create(jobCreation)

	if errors.Is(db.Error, gorm.ErrInvalidData) {
		return &models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      db.Error,
		}
	}
	return nil
}

// Admin updates their job posts
func (database AdminRepo) UpdateJobPost(jobCreation *models.JobCreation, jobID int) *models.ErrorResponse {
	var newJobCreation models.JobCreation

	db := database.
		Where("job_id = ?", jobID).
		First(&newJobCreation)

	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      db.Error,
		}
	}

	updatDb := database.Model(&models.JobCreation{}).Where("job_id = ?", jobID).Updates(map[string]interface{}{
		"JobStatus": jobCreation.JobStatus,
		"Vacancy":   jobCreation.Vacancy,
	})

	if errors.Is(updatDb.Error, gorm.ErrInvalidData) {
		return &models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      db.Error,
		}
	} else if updatDb.RowsAffected == 0 {
		return &models.ErrorResponse{
			StatusCode: http.StatusBadGateway,
			Error:      fmt.Errorf("no rows affected "),
		}
	}

	return nil
}

// Admin Get their job applied details(user) by role
func (database AdminRepo) GetPostByRole(jobRole string, userID int) ([]models.UserJobDetails, *models.ErrorResponse) {
	var userJobDetails []models.UserJobDetails

	data := database.
		Where("job_role=?", jobRole).
		First(&userJobDetails)

	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("no applications found for this Job Role"),
		}
	}

	dbvalue := database.Preload("User").
		Where("job_role=? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ?)", jobRole, userID).Find(&userJobDetails)

	if errors.Is(dbvalue.Error, gorm.ErrRecordNotFound) {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("cant'able to find your details Properly,Give him correctly"),
		}
	} else if len(userJobDetails) == 0 {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      fmt.Errorf("not have access to view this details"),
		}
	}

	return userJobDetails, nil
}

// Admin get their job applied details by ID
func (database AdminRepo) GetPostByJobID(jobID int, userID int) ([]models.UserJobDetails, *models.ErrorResponse) {
	var userJobdetails []models.UserJobDetails

	result := database.
		Where("job_id=?", jobID).
		First(&userJobdetails)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("no applications found for this Job ID"),
		}
	}

	dbvalue := database.Preload("User").
		Where("job_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ? )", jobID, userID).
		Find(&userJobdetails)

	if errors.Is(dbvalue.Error, gorm.ErrRecordNotFound) {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      dbvalue.Error,
		}
	} else if len(userJobdetails) == 0 {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusUnauthorized,
			Error:      fmt.Errorf("not authorized to view this Job ID"),
		}
	}

	return userJobdetails, nil
}

// Admin get Job applied details by USER ID
func (database AdminRepo) GetPostByUserID(applicantID int, userID int) ([]models.UserJobDetails, *models.ErrorResponse) {
	var userJobdetails []models.UserJobDetails

	data := database.Debug().Preload("Job").
		Where("user_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ?)", applicantID, userID).
		Find(&userJobdetails)

	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("failed to fetch job details: %s", data.Error),
		}
	} else if len(userJobdetails) == 0 {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("no job details found for the given user"),
		}
	}
	return userJobdetails, nil
}

// Admin get their Own Post details
func (database AdminRepo) GetPosts(userID int) ([]models.JobCreation, *models.ErrorResponse) {
	var jobCreation []models.JobCreation

	dbvalue := database.
		Where("admin_id = ?", userID).
		Find(&jobCreation)

	if errors.Is(dbvalue.Error, gorm.ErrRecordNotFound) {
		return nil, &models.ErrorResponse{
			Error:      dbvalue.Error,
			StatusCode: http.StatusBadRequest,
		}
	} else if len(jobCreation) == 0 {
		return nil, &models.ErrorResponse{
			Error:      fmt.Errorf("no details found for this userID"),
			StatusCode: http.StatusNotFound,
		}
	}
	return jobCreation, nil
}
