package repository

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type IAdminRepo interface {
	CreateJobPost(jobCreation *models.JobCreation) *dto.ErrorResponse
	GetApplicantAndJobDetails(jobDetails map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64)
	GetJobsAppliedByUser(userIDJobs map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64)
	GetJobsCreated(createdPosts map[string]interface{}) ([]models.JobCreation, *dto.ErrorResponse, int64)
	UpdateJobPost(jobData *models.JobCreation, jobID int, userID int) *dto.ErrorResponse
	DeleteJobPost() *dto.ErrorResponse
}
type AdminRepo struct {
	*internals.NewConnection
}

func InitAdminRepo(db *internals.NewConnection) IAdminRepo {
	return &AdminRepo{
		db,
	}
}

// Admin creates new post in this case
func (database *AdminRepo) CreateJobPost(jobCreation *models.JobCreation) *dto.ErrorResponse {
	db := database.Create(jobCreation)
	if errors.Is(db.Error, gorm.ErrInvalidData) {
		return &dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      fmt.Errorf("can't able to create the jobPost"),
		}
	} else if db.Error != nil {
		return &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      db.Error,
		}
	}

	return nil
}

func (database *AdminRepo) GetApplicantAndJobDetails(jobDetails map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64) {
	var userJobDetails []models.UserJobDetails
	var count int64

	jobRole := jobDetails["jobRole"].(string)
	jobID := jobDetails["jobID"].(int)
	userID := jobDetails["userID"].(int)
	limit := jobDetails["limit"].(int)
	offset := jobDetails["offset"].(int)

	query := database.Debug()

	if jobRole != "" {
		result := database.
			Where(&models.UserJobDetails{JobRole: jobRole}).
			First(&userJobDetails)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil, &dto.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the job applications for Job Role"),
				}, 0
			}
			return nil, &dto.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      result.Error,
			}, 0
		}

		query = query.Preload("User").
			Where("job_role=? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ?)", jobRole, userID).
			Find(&userJobDetails)
	}

	if jobID != 0 {
		result := database.
			Where("job_id=?", jobID).
			First(&userJobDetails)
		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil, &dto.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the job applications for Job ID"),
				}, 0
			}
			return nil, &dto.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      result.Error,
			}, 0
		}

		query = query.Preload("User").
			Where("job_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ? )", jobID, userID).
			Find(&userJobDetails)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Errorf("failed to count job details: %v", err),
		}, 0
	}

	result := query.Limit(limit).Offset(offset).Find(&userJobDetails)
	if result.Error != nil {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      result.Error,
		}, 0
	} else if len(userJobDetails) == 0 {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("your not access to view this Details"),
		}, 0
	}

	return userJobDetails, nil, count
}

// Admin get Job applied details by USER ID
func (database *AdminRepo) GetJobsAppliedByUser(userIDJobs map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64) {
	var userJobdetails []models.UserJobDetails
	var count int64

	limit := userIDJobs["limit"].(int)
	offset := userIDJobs["offset"].(int)
	applicantID := userIDJobs["applicant"].(int)
	userID := userIDJobs["userID"].(int)

	query := database.Model(&models.UserJobDetails{}).
		Where("user_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ?)", applicantID, userID)
	if err := query.Count(&count).Error; err != nil {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Errorf("failed to count job details: %v", err),
		}, 0
	}

	data := query.Preload("Job").
		Limit(limit).Offset(offset).
		Find(&userJobdetails)
	if errors.Is(data.Error, gorm.ErrRecordNotFound) {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("failed to fetch job details: %s", data.Error),
		}, 0
	}

	if len(userJobdetails) == 0 {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("no job details found for the given user"),
		}, 0
	}

	return userJobdetails, nil, count
}

// Admin get their Own Post details
func (database *AdminRepo) GetJobsCreated(searchMaps map[string]interface{}) ([]models.JobCreation, *dto.ErrorResponse, int64) {
	var jobCreation []models.JobCreation
	var count int64

	userID := searchMaps["userID"].(int)
	limit := searchMaps["limit"].(int)
	offset := searchMaps["offset"].(int)
	jobRole := searchMaps["jobRole"].(string)
	country := searchMaps["country"].(string)

	query := database.Debug()

	if jobRole != "" {
		dbRole := database.Where(&models.JobCreation{JobRole: jobRole}).First(&models.JobCreation{})
		if dbRole.Error != nil {
			if errors.Is(dbRole.Error, gorm.ErrRecordNotFound) {
				return nil, &dto.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the specified Jobrole, please check it"),
				}, 0
			}
			return nil, &dto.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      dbRole.Error,
			}, 0
		}

		query = query.Debug().Where(&models.JobCreation{JobRole: jobRole})
	}

	if country != "" {
		dbCountry := database.Where(&models.JobCreation{Country: country}).First(&models.JobCreation{})
		if dbCountry.Error != nil {
			if errors.Is(dbCountry.Error, gorm.ErrRecordNotFound) {
				return nil, &dto.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the specified Country, please check it"),
				}, 0
			}
			return nil, &dto.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      dbCountry.Error,
			}, 0
		}

		query = query.Debug().Where(&models.JobCreation{Country: country})
	}

	if userID < 0 {
		dbUserId := database.Where(&models.JobCreation{AdminID: userID}).First(&models.JobCreation{})
		if dbUserId.Error != nil {
			if errors.Is(dbUserId.Error, gorm.ErrRecordNotFound) {
				return nil, &dto.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the specified Country, please check it"),
				}, 0
			}
			return nil, &dto.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      dbUserId.Error,
			}, 0
		}

		query = query.Debug().Where(&models.JobCreation{AdminID: userID})
	}

	result := query.Debug().Limit(limit).Offset(offset).Find(&jobCreation)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, &dto.ErrorResponse{
			Error:      result.Error,
			StatusCode: http.StatusBadRequest,
		}, 0
	}

	return jobCreation, nil, count
}

// Admin updates their job posts
func (database *AdminRepo) UpdateJobPost(jobData *models.JobCreation, jobID int, userID int) *dto.ErrorResponse {
	var jobUpdate models.JobCreation

	db := database.
		Where("job_id = ?", jobID).
		First(&jobUpdate)
	if errors.Is(db.Error, gorm.ErrRecordNotFound) {
		return &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      db.Error,
		}
	}

	updatDb := database.Model(&models.JobCreation{}).Where("job_id = ?", jobID).Updates(map[string]interface{}{
		"JobStatus": jobData.JobStatus,
		"Vacancy":   jobData.Vacancy,
	})
	if errors.Is(updatDb.Error, gorm.ErrInvalidData) {
		return &dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      db.Error,
		}
	} else if updatDb.RowsAffected == 0 {
		return &dto.ErrorResponse{
			StatusCode: http.StatusBadGateway,
			Error:      fmt.Errorf("can't able to update jobPost"),
		}
	}

	return nil
}

func (database *AdminRepo) DeleteJobPost() *dto.ErrorResponse {
	var jobDeletion []models.JobCreation

	now := time.Now()
	lastYear := now.AddDate(-1, 0, 0)

	result := database.Debug().Select("created_at").Where("created_at <= ?", lastYear).Find(&jobDeletion).Delete(&jobDeletion)
	if result.Error != nil {
		return &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      result.Error,
		}
	} else if result.RowsAffected == 0 {
		return &dto.ErrorResponse{
			StatusCode: http.StatusNotModified,
			Error:      fmt.Errorf("no data are deleted here"),
		}
	}

	return nil
}
