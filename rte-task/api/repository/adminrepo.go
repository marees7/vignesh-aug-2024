package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	CreatePostDetailsByAdmin(user *models.JobCreation) error
	UpdateJobPostsByAdmin(user *models.JobCreation, jobID int, adminID int) error
	GetDetailsByRoleByAdmin(user *[]models.UserJobDetails, jobrole string, adminid int) error
	GetJobDetailsByJobIdByAdmin(user *[]models.UserJobDetails, jobID int, adminid int) error
	GetJobDetailsByUserIdByAdmin(user *[]models.UserJobDetails, userId int, adminvalues int) error
	GetOwnPostDetailsByAdmin(user *[]models.JobCreation, adminid int) error
}
type adminRepo struct {
	*gorm.DB
}

// Admin creates new post in this case
func (database *adminRepo) CreatePostDetailsByAdmin(user *models.JobCreation) error {
	dbvalues := database.Create(user)
	if dbvalues.Error != nil {
		return fmt.Errorf("can't able to create post details")
	}

	return nil
}

// Admin updates their job posts
func (database *adminRepo) UpdateJobPostsByAdmin(user *models.JobCreation, jobID int, adminID int) error {
	var job models.JobCreation

	result := database.Where("job_id = ?", jobID).First(&job)
	if result.Error != nil {
		return fmt.Errorf("cannot find the job with JobID: %d, please provide correct details", jobID)
	}

	if job.DomainID != adminID {
		return fmt.Errorf("you are not authorized to update this job post")
	}

	updateResult := database.Model(&models.JobCreation{}).Where("job_id = ?", jobID).Updates(map[string]interface{}{
		"JobStatus": user.JobStatus,
		"Vacancy":   user.Vacancy,
	})

	if updateResult.Error != nil {
		return fmt.Errorf("unable to update the job post")
	}

	return nil
}

// Admin Get their job applied details(user) by role
func (database *adminRepo) GetDetailsByRoleByAdmin(user *[]models.UserJobDetails, roletype string, adminid int) error {

	data := database.Where("job_role=?", roletype).First(&user)
	if data.Error != nil {
		return fmt.Errorf("no applications found for this job role")
	}

	dbvalue := database.Preload("User").
		Where("job_role=? AND job_id IN (SELECT job_id FROM job_creations WHERE domain_id = ?)", roletype, adminid).Find(&user)
	if dbvalue.Error != nil {
		return fmt.Errorf("cant'able to find your details Properly,Give him correctly")
	}
	if dbvalue.RowsAffected == 0 {
		return fmt.Errorf("not have access to view this details")
	}

	return nil
}

// Admin get their job applied details by ID
func (database *adminRepo) GetJobDetailsByJobIdByAdmin(user *[]models.UserJobDetails, jobID int, adminid int) error {
	data := database.Where("job_id=?", jobID).First(&user)
	if data.Error != nil {
		return fmt.Errorf("no applications found for this job role")
	}

	dbvalue := database.Preload("User").
		Where("job_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE domain_id = ? )", jobID, adminid).
		Find(&user)

	if dbvalue.Error != nil {
		return fmt.Errorf("cant'able to find your Details Properly,Give him correctly")
	}
	if dbvalue.RowsAffected == 0 {
		return fmt.Errorf("not have access to view this details")
	}

	return nil
}

// Admin get Job applied details by USER ID
func (database *adminRepo) GetJobDetailsByUserIdByAdmin(user *[]models.UserJobDetails, userId int, adminid int) error {
	result := database.Preload("Job").
		Where("user_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE domain_id = ?)", userId, adminid).
		Find(&user)

	if result.Error != nil {
		return fmt.Errorf("failed to fetch job details: %s", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no job details found for the given user")
	}

	return nil
}

// Admin get their Own Post details
func (database *adminRepo) GetOwnPostDetailsByAdmin(user *[]models.JobCreation, adminid int) error {
	data := database.Where("domain_id=?", adminid).First(&user)
	if data.Error != nil {
		return fmt.Errorf("can't able to find any job posts by this admin")
	}
	dbvalue := database.
		Where("domain_id = ?", adminid).
		Find(&user)

	if dbvalue.Error != nil {
		return fmt.Errorf("cant'able to find your Details Properly,Give him correctly")
	}
	
	return nil
}
