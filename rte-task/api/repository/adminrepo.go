package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type I_AdminRepo interface {
	CreatePostForUser(user *models.JobCreation) error
	UpdateJobPost(user *models.JobCreation, jobID int) error
	GetByRoles(roleType string, roleID int) ([]models.UserJobDetails, error)
	GetByJobID(jobid int, roleID int) ([]models.UserJobDetails, error)
	GetByUserID(userID int, roleID int) ([]models.UserJobDetails, error)
	GetAllOwnPosts(roleID int) ([]models.JobCreation, error)
}
type AdminRepo struct {
	*internals.ConnectionNew
}

func GetAdminRepository(db *internals.ConnectionNew) I_AdminRepo {
	return &AdminRepo{
		db,
	}
}

// Admin creates new post in this case
func (database *AdminRepo) CreatePostForUser(user *models.JobCreation) error {
	dbvalues := database.Create(user)
	if dbvalues.Error != nil {
		return fmt.Errorf("can't able to create post details")
	}

	return nil
}

// Admin updates their job posts
func (database *AdminRepo) UpdateJobPost(user *models.JobCreation, jobID int) error {
	var job models.JobCreation

	result := database.Where("job_id = ?", jobID).First(&job)
	if result.Error != nil {
		return fmt.Errorf("cannot find the job with JobID: %d, please provide correct details", jobID)
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
func (database *AdminRepo) GetByRoles(roleType string, roleID int) ([]models.UserJobDetails, error) {
	var user []models.UserJobDetails
	data := database.Where("job_role=?", roleType).First(&user)
	if data.Error != nil {
		return nil, fmt.Errorf("no applications found for this job role")
	}

	dbvalue := database.Preload("User").
		Where("job_role=? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ?)", roleType, roleID).Find(&user)
	if dbvalue.Error != nil {
		return nil, fmt.Errorf("cant'able to find your details Properly,Give him correctly")
	}
	if dbvalue.RowsAffected == 0 {
		return nil, fmt.Errorf("not have access to view this details")
	}

	return user, nil
}

// Admin get their job applied details by ID
func (database *AdminRepo) GetByJobID(jobID int, roleID int) ([]models.UserJobDetails, error) {
	var user []models.UserJobDetails
	data := database.Where("job_id=?", jobID).First(&user)
	if data.Error != nil {
		return nil, fmt.Errorf("no applications found for this job role")
	}

	dbvalue := database.Preload("User").
		Where("job_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ? )", jobID, roleID).
		Find(&user)

	if dbvalue.Error != nil {
		return nil, fmt.Errorf("cant'able to find your Details Properly,Give him correctly")
	}
	if dbvalue.RowsAffected == 0 {
		return nil, fmt.Errorf("not have access to view this details")
	}

	return user, nil
}

// Admin get Job applied details by USER ID
func (database *AdminRepo) GetByUserID(userID int, roleID int) ([]models.UserJobDetails, error) {
	var user []models.UserJobDetails
	result := database.Preload("Job").
		Where("user_id = ? AND job_id IN (SELECT job_id FROM job_creations WHERE admin_id = ?)", userID, roleID).
		Find(&user)

	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch job details: %s", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("no job details found for the given user")
	}

	return user, nil
}

// Admin get their Own Post details
func (database *AdminRepo) GetAllOwnPosts(roleID int) ([]models.JobCreation, error) {
	var user []models.JobCreation
	data := database.Where("admin_id=?", roleID).First(&user)
	if data.Error != nil {
		return nil, fmt.Errorf("can't able to find any job posts by this admin")
	}
	dbvalue := database.
		Where("admin_id = ?", roleID).
		Find(&user)

	if dbvalue.Error != nil {
		return nil, fmt.Errorf("cant'able to find your Details Properly,Give him correctly")
	}

	return user, nil
}
