package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type I_UserRepo interface {
	GetAllPost() ([]models.JobCreation, error)
	GetPostsByRoles(jobrole string, country string) ([]models.JobCreation, error)
	GetByCompanys(company string) ([]models.JobCreation, error)
	CreateJobApplication(user *models.UserJobDetails) error
	GetUserJob(roleID int) ([]models.UserJobDetails, error)
	GetUserId(user *models.UserJobDetails) error
}

type Userrepo struct {
	*internals.ConnectionNew
}

func GetUserRepository(db *internals.ConnectionNew) I_UserRepo {
	return &Userrepo{
		db,
	}
}

// get their all post details by admin or users
func (database *Userrepo) GetAllPost() ([]models.JobCreation, error) {
	var user []models.JobCreation
	dbvalues := database.Find(&user)
	if dbvalues.Error != nil {
		return nil, fmt.Errorf("error occured in while getting the post details")
	}

	return user, nil
}

// retrive thier job details by their JobRole by users or admin
func (database *Userrepo) GetPostsByRoles(jobrole string, country string) ([]models.JobCreation, error) {
	var user []models.JobCreation
	data := database.Where("job_title=?", jobrole).First(&user)
	if data.Error != nil {
		return nil, fmt.Errorf("no one can post the job posts based on your role and country")
	}
	dbvalue := database.Where(&models.JobCreation{Country: country}).First(&user)
	if dbvalue.Error != nil {
		return nil, fmt.Errorf("cant'able to find your country Properly,Give him correctly")
	}
	dbValue := database.Where(&models.JobCreation{JobTitle: jobrole, Country: country}).Find(&user)
	if dbValue.Error != nil {
		return nil, fmt.Errorf("cant'able to find your details properly,Give him correctly")
	}

	return user, nil
}

// get by thier Company Names by particular Details by users or admin
func (database *Userrepo) GetByCompanys(company string) ([]models.JobCreation, error) {
	var user []models.JobCreation
	data := database.Where("company_name=?", company).First(&user)
	if data.Error != nil {
		return nil, fmt.Errorf("no one can post the job for this company Name")
	}

	dbValue := database.Where(&models.JobCreation{CompanyName: company}).Find(&user)
	if dbValue.Error != nil {
		return nil, fmt.Errorf("cant'able to find your jobs in that company,Give him correctly")
	}

	return user, nil
}

// users apply for the job post
func (database *Userrepo) CreateJobApplication(user *models.UserJobDetails) error {
	dbvalues := database.Create(user)
	if dbvalues.Error != nil {
		return fmt.Errorf("can't able to apply the job post here,please Check")
	}
	return nil
}

// users get thier applied details by their own Ids
func (database *Userrepo) GetUserJob(roleID int) ([]models.UserJobDetails, error) {
	var user []models.UserJobDetails
	data := database.Where("user_id=?", roleID).First(&user)
	if data.Error != nil {
		return nil, fmt.Errorf("cant'able to find your UserId ,Give him correctly")
	}

	dbvalue := database.
		Preload("Job").
		Where("user_id = ?", roleID).
		Find(&user)

	if dbvalue.Error != nil {
		return nil, fmt.Errorf("cant'able to create your details ,Give him correctly")
	}
	return user, nil
}

// Check if user ID is applied for the Job or Not
func (database *Userrepo) GetUserId(user *models.UserJobDetails) error {
	var count int64
	var newuser *models.JobCreation
	
	jobid := database.Model(&models.JobCreation{}).Where("job_id=? ", user.JobID).First(&newuser)
	if jobid.Error != nil {
		return fmt.Errorf("unable to fetch Job Details properly,Check it JobId once")
	}
	if newuser.JobStatus == "COMPLETED" {
		return fmt.Errorf("this Job Application is closed")
	}

	data := database.Model(&models.UserJobDetails{}).Where("user_id=? AND job_id=?", user.UserId, user.JobID).Count(&count)
	if count > 0 {
		return fmt.Errorf("already registered,You have applied for this job")
	}
	if data.Error != nil {
		return fmt.Errorf("error occured,While applying the post")
	}
	return nil
}
