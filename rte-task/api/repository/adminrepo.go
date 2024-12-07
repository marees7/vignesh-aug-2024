package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	RepoFindAllUsers(user *[]models.UserDetails) error
	RepoCreateNewPost(user *models.JobCreation) error
	RepoUpdateJobPost(user *models.JobCreation, jobID int, adminID int) error
	RepoGetJobAppliedDetailsbyrole(user *[]models.UserJobDetails, jobrole string) error
	RepoGetJobAppliedDetailsByJobId(user *[]models.UserJobDetails, jobID int) error
	RepoGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int, adminvalues int) error
	RepoGetPostedDetailsByAdmin(user *[]models.JobCreation, adminid int) error
}
type adminRepo struct {
	*gorm.DB
}

// Find their All users
func (database *adminRepo) RepoFindAllUsers(user *[]models.UserDetails) error {
	value := database.Where("role_type=?", "USER").Find(&user)
	if value.Error != nil {
		return fmt.Errorf("cant'able to find your users Properly,Give him correctly")
	}
	return nil
}

// Admin creates new post in this case
func (database *adminRepo) RepoCreateNewPost(user *models.JobCreation) error {
	dbvalues := database.Create(user)
	if dbvalues.Error != nil {
		return fmt.Errorf("can't able to create post details")
	}
	return nil
}

// Admin updates their job posts
func (database *adminRepo) RepoUpdateJobPost(user *models.JobCreation, jobID int, adminID int) error {
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
func (database *adminRepo) RepoGetJobAppliedDetailsbyrole(user *[]models.UserJobDetails, roletype string) error {

	data := database.Where("job_role=?", roletype).First(&user)
	if data.Error != nil {
		return fmt.Errorf("no one can apply for this job ")
	}

	dbvalue := database.Preload("User").
		Where(&models.UserJobDetails{JobRole: roletype}).Find(&user)
	if dbvalue.Error != nil {
		return fmt.Errorf("cant'able to find your details Properly,Give him correctly")
	}
	return nil
}

// Admin get their job applied details by ID
func (database *adminRepo) RepoGetJobAppliedDetailsByJobId(user *[]models.UserJobDetails, jobID int) error {
	data := database.Where("job_id=?", jobID).First(&user)
	if data.Error != nil {
		return fmt.Errorf("no one can apply for this job")
	}

	dbvalue := database.Preload("User").
		Where("job_id = ?", jobID).
		Find(&user)

	if dbvalue.Error != nil {
		return fmt.Errorf("cant'able to find your Details Properly,Give him correctly")
	}
	return nil
}

// Admin get Job applied details by USER ID
func (database *adminRepo) RepoGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int, adminid int) error {
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
func (database *adminRepo) RepoGetPostedDetailsByAdmin(user *[]models.JobCreation, adminid int) error {
	data := database.Where("domain_id=?", adminid).First(&user)
	if data.Error != nil {
		return fmt.Errorf("no one can apply for this job")
	}
	dbvalue := database.
		Where("domain_id = ?", adminid).
		Find(&user)

	if dbvalue.Error != nil {
		return fmt.Errorf("cant'able to find your Details Properly,Give him correctly")
	}
	return nil
}
