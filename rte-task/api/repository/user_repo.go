package repository

import (
	"fmt"
	"net/http"

	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type IUserRepo interface {
	GetAllJobPosts(company string) ([]models.JobCreation, *models.ErrorResponse)
	GetJobByRole(jobrole string, country string) ([]models.JobCreation, *models.ErrorResponse)
	CreateApplication(userJobDetails *models.UserJobDetails) *models.ErrorResponse
	GetUserAppliedJobs(roleID int) ([]models.UserJobDetails, *models.ErrorResponse)
	GetUserID(user *models.UserJobDetails) *models.ErrorResponse
}

type Userrepo struct {
	*internals.ConnectionNew
}

func InitUserRepo(db *internals.ConnectionNew) IUserRepo {
	return Userrepo{
		db,
	}
}

// get their all post details by admin or users
func (database Userrepo) GetAllJobPosts(company string) ([]models.JobCreation, *models.ErrorResponse) {
	var jobCreation []models.JobCreation
	if company == "" {
		db := database.Find(&jobCreation)
		if db.Error != nil {
			return nil, &models.ErrorResponse{
				StatusCode: http.StatusNotFound,
				Error:      fmt.Errorf("error occured in while getting the post details"),
			}
		}
		return jobCreation, nil
	} else {
		data := database.
			Where("company_name=?", company).
			First(&jobCreation)

		if data.Error != nil {
			return nil, &models.ErrorResponse{
				StatusCode: http.StatusNotFound,
				Error:      fmt.Errorf("no one can post the job for in this Company"),
			}
		}

		datas := database.Where(&models.JobCreation{CompanyName: company}).Find(&jobCreation)
		if datas.Error != nil {
			return nil, &models.ErrorResponse{
				StatusCode: http.StatusNotFound,
				Error:      fmt.Errorf("cant'able to find your jobs in that company,Give him correctly"),
			}
		}
		return jobCreation, nil
	}
}

// retrive thier job details by their JobRole by users or admin
func (database Userrepo) GetJobByRole(jobrole string, country string) ([]models.JobCreation, *models.ErrorResponse) {
	var jobCreation []models.JobCreation

	db := database.
		Where("job_title=?", jobrole).
		First(&jobCreation)

	if db.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("cant'able to find your Jobrole Properly,Check it once"),
		}
	}

	data := database.
		Where(&models.JobCreation{Country: country}).
		First(&jobCreation)

	if data.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("cant'able to find your country Properly,Check it "),
		}
	}

	datas := database.
		Where(&models.JobCreation{JobTitle: jobrole, Country: country}).
		Find(&jobCreation)

	if datas.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      fmt.Errorf("cant'able to find your details properly,Give him correctly"),
		}
	}
	return jobCreation, nil
}

// users apply for the job post
func (database Userrepo) CreateApplication(user *models.UserJobDetails) *models.ErrorResponse {
	db := database.Create(user)
	if db.Error != nil {
		return &models.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      fmt.Errorf("can't able to apply the job post here,please Check"),
		}
	}
	return nil
}

// users get thier applied details by their own Ids
func (database Userrepo) GetUserAppliedJobs(roleID int) ([]models.UserJobDetails, *models.ErrorResponse) {
	var userJobDetails []models.UserJobDetails
	db := database.
		Where("user_id=?", roleID).
		First(&userJobDetails)

	if db.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("cant'able to find your UserId ,Give him correctly"),
		}
	}

	data := database.
		Preload("Job").
		Where("user_id = ?", roleID).
		Find(&userJobDetails)

	if data.Error != nil {
		return nil, &models.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      fmt.Errorf("cant'able to create your details ,Give him correctly"),
		}
	}
	return userJobDetails, nil
}

// Check if user ID is applied for the Job or Not
func (database Userrepo) GetUserID(user *models.UserJobDetails) *models.ErrorResponse {
	var count int64
	var jobCreation *models.JobCreation

	jobID := database.
		Model(&models.JobCreation{}).
		Where("job_id=? ", user.JobID).
		First(&jobCreation)

	if jobID.Error != nil {
		return &models.ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      fmt.Errorf("unable to fetch Job Details properly,Check it JobId once"),
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
		Where("user_id=? AND job_id=?", user.UserId, user.JobID).
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
