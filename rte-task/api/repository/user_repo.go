package repository

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type IUserRepo interface {
	CreateApplication(userJobDetails *models.UserJobDetails) *models.ErrorResponse
	GetAllJobPosts(searchJobs map[string]interface{}) ([]models.JobCreation, *models.ErrorResponse, int64)
	GetUserAppliedJobs(userJobs map[string]interface{}) ([]models.UserJobDetails, *models.ErrorResponse, int64)
	GetUserByID(userJobDetails *models.UserJobDetails) *models.ErrorResponse
}

type Userrepo struct {
	*internals.ConnectionNew
}

func InitUserRepo(db *internals.ConnectionNew) IUserRepo {
	return &Userrepo{
		db,
	}
}

// users apply for the job post
func (database Userrepo) CreateApplication(userJobDetails *models.UserJobDetails) *models.ErrorResponse {
	db := database.Debug().Create(&userJobDetails)

	if db.Error != nil {
		return &models.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      db.Error,
		}
	}

	return nil
}

// get their all post details by admin or users
func (database *Userrepo) GetAllJobPosts(searchJobs map[string]interface{}) ([]models.JobCreation, *models.ErrorResponse, int64) {
	var jobCreation []models.JobCreation
	var count int64

	limit := searchJobs["Limit"].(int)
	offset := searchJobs["Offset"].(int)
	jobRole := searchJobs["Role"].(string)
	country := searchJobs["Country"].(string)
	company := searchJobs["Company"].(string)

	query := database.Debug()

	if company != "" {
		result := database.Where(&models.JobCreation{CompanyName: company}).First(&jobCreation)

		if result.Error != nil {
			if errors.Is(result.Error, gorm.ErrRecordNotFound) {
				return nil, &models.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the specified company, please check it"),
				}, 0
			}
			return nil, &models.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      result.Error,
			}, 0
		}

		query = query.Where(&models.JobCreation{CompanyName: company}).Find(&jobCreation)
	}

	if jobRole != "" {
		dbRole := database.Where(&models.JobCreation{JobRole: jobRole}).First(&models.JobCreation{})

		if dbRole.Error != nil {
			if errors.Is(dbRole.Error, gorm.ErrRecordNotFound) {
				return nil, &models.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the specified Jobrole, please check it"),
				}, 0
			}
			return nil, &models.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      dbRole.Error,
			}, 0
		}

		query = query.Where(&models.JobCreation{JobRole: jobRole}).Find(&jobCreation)
	}

	if country != "" {
		dbCountry := database.Where(&models.JobCreation{Country: country}).First(&models.JobCreation{})

		if dbCountry.Error != nil {
			if errors.Is(dbCountry.Error, gorm.ErrRecordNotFound) {
				return nil, &models.ErrorResponse{
					StatusCode: http.StatusNotFound,
					Error:      fmt.Errorf("unable to fetch the specified Country, please check it"),
				}, 0
			}
			return nil, &models.ErrorResponse{
				StatusCode: http.StatusInternalServerError,
				Error:      dbCountry.Error,
			}, 0
		}

		query = query.Where(&models.JobCreation{Country: country}).Find(&jobCreation)
	}

	result := query.Debug().Limit(limit).Offset(offset).Find(&jobCreation).Count(&count)

	if result.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      result.Error,
		}, 0
	} else if len(jobCreation) == 0 {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("unable to fetch Job Details properly, based on your Criteria"),
		}, 0
	}

	return jobCreation, nil, count
}

// users get thier applied details by their own Ids
func (database *Userrepo) GetUserAppliedJobs(userJobs map[string]interface{}) ([]models.UserJobDetails, *models.ErrorResponse, int64) {
	var userJobDetails []models.UserJobDetails
	var count int64

	limit := userJobs["Limit"].(int)
	offset := userJobs["Offset"].(int)
	userID := userJobs["UserID"].(int)

	db := database.
		Where("user_id=?", userID).
		First(&userJobDetails)

	if db.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("unable to fetch the User Applied Jobs properly,Check it UserId once"),
		}, 0
	}

	data := database.
		Preload("Job").
		Where("user_id = ?", userID).Limit(limit).Offset(offset).
		Find(&userJobDetails).Count(&count)

	if data.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      fmt.Errorf("cant'able to create your details ,Give him correctly"),
		}, 0
	}

	return userJobDetails, nil, count
}

// Check if user ID is applied for the Job or Not
func (database *Userrepo) GetUserByID(userJobDetails *models.UserJobDetails) *models.ErrorResponse {
	var count int64
	var jobCreation *models.JobCreation

	jobID := database.
		Model(&models.JobCreation{}).
		Where("job_id=? ", userJobDetails.JobID).
		First(&jobCreation)

	if jobID.Error != nil {
		return &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("unable to fetch User Details,Check it UserID once"),
		}
	}

	if jobCreation.JobStatus == "COMPLETED" {
		return &models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      fmt.Errorf("this Job Application is closed"),
		}
	}

	db := database.
		Model(&models.UserJobDetails{}).
		Where("user_id=? AND job_id=?", userJobDetails.UserID, userJobDetails.JobID).
		Count(&count)

	if count > 0 {
		return &models.ErrorResponse{
			StatusCode: http.StatusAlreadyReported,
			Error:      fmt.Errorf("already registered,You have applied for this job"),
		}
	}

	if db.Error != nil {
		return &models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      db.Error,
		}
	}

	return nil
}
