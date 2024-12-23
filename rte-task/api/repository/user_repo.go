package repository

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	CreateApplication(userJobDetails *models.UserJobDetails) *dto.ErrorResponse
	GetAllJobPosts(searchJobs map[string]interface{}) ([]models.JobCreation, *dto.ErrorResponse, int64)
	GetUserAppliedJobs(userJobs map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64)
	GetUserByID(userJobDetails *models.UserJobDetails) *dto.ErrorResponse
}

type Userrepo struct {
	*internals.NewConnection
}

func InitUserRepo(db *internals.NewConnection) IUserRepo {
	return &Userrepo{
		db,
	}
}

// users apply for the job post
func (database Userrepo) CreateApplication(userJobDetails *models.UserJobDetails) *dto.ErrorResponse {
	db := database.Create(&userJobDetails)

	if errors.Is(db.Error, gorm.ErrInvalidData) {
		return &dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      fmt.Errorf("can't able to create the apllication"),
		}
	} else if db.Error != nil {
		return &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      db.Error,
		}
	}

	return nil
}

// get their all post details by admin or users
func (database *Userrepo) GetAllJobPosts(searchJobs map[string]interface{}) ([]models.JobCreation, *dto.ErrorResponse, int64) {
	var jobCreation []models.JobCreation
	var count int64

	limit := searchJobs["limit"].(int)
	offset := searchJobs["offset"].(int)
	jobRole := searchJobs["role"].(string)
	country := searchJobs["country"].(string)
	company := searchJobs["company"].(string)

	query := database.Model(&models.JobCreation{})

	if company != "" {
		query = query.Where("company_name = ?", company)
	}
	if jobRole != "" {
		query = query.Where("job_role = ?", jobRole)
	}
	if country != "" {
		query = query.Where("country = ?", country)
	}

	if err := query.Count(&count).Error; err != nil {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Errorf("unable to count records: %v", err),
		}, 0
	}

	if err := query.Limit(limit).Offset(offset).Find(&jobCreation).Error; err != nil {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Errorf("unable to fetch records: %v", err),
		}, 0
	}

	if len(jobCreation) == 0 {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("no job details found based on your criteria"),
		}, 0
	}

	return jobCreation, nil, count
}

// users get thier applied details by their own Ids
func (database *Userrepo) GetUserAppliedJobs(userJobs map[string]interface{}) ([]models.UserJobDetails, *dto.ErrorResponse, int64) {
	var userJobDetails []models.UserJobDetails
	var count int64

	limit := userJobs["limit"].(int)
	offset := userJobs["offset"].(int)
	userID := userJobs["userID"].(int)

	db := database.
		Where("user_id = ?", userID).
		Model(&models.UserJobDetails{}).
		Count(&count)
	if db.Error != nil {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("unable to fetch the User Applied Jobs properly, check the UserID once"),
		}, 0
	}

	db = database.
		Preload("Job").
		Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&userJobDetails)
	if db.Error != nil {
		return nil, &dto.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Errorf("unable to get your details, please provide them correctly"),
		}, 0
	}

	return userJobDetails, nil, count
}

// Check if user ID is applied for the Job or Not
func (database *Userrepo) GetUserByID(userJobDetails *models.UserJobDetails) *dto.ErrorResponse {
	var count int64
	var jobCreation *models.JobCreation

	jobID := database.
		Model(&models.JobCreation{}).
		Where("job_id=? ", userJobDetails.JobID).
		First(&jobCreation)
	if jobID.Error != nil {
		return &dto.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("unable to fetch User Details,Check it UserID once"),
		}
	}

	if jobCreation.JobStatus == "COMPLETED" {
		return &dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      fmt.Errorf("this Job Application is closed"),
		}
	}

	db := database.
		Model(&models.UserJobDetails{}).
		Where("user_id=? AND job_id=?", userJobDetails.UserID, userJobDetails.JobID).
		Count(&count)

	if count > 0 {
		return &dto.ErrorResponse{
			StatusCode: http.StatusAlreadyReported,
			Error:      fmt.Errorf("already registered,You have applied for this job"),
		}
	}

	if db.Error != nil {
		return &dto.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      db.Error,
		}
	}

	return nil
}
