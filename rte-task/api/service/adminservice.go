package service

import (
	"github.com/Vigneshwartt/golang-rte-task/api/repository"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type I_AdminService interface {
	CreateJobPosts(user *models.JobCreation) error
	UpdatePosts(user *models.JobCreation, jobid int) error
	GetAllPostsByRole(roleType string, roleID int) ([]models.UserJobDetails, error)
	GetAllPostsByJobID(jobid int, roleID int) ([]models.UserJobDetails, error)
	GetAllPostsByUserID(userID int, roleID int) ([]models.UserJobDetails, error)
	GetOwnPost(roleID int) ([]models.JobCreation, error)
}
type AdminService struct {
	repository.I_AdminRepo
}

func GetAdminService(db repository.I_AdminRepo) I_AdminService {
	return &AdminService{
		db,
	}
}

// create their Jobposts
func (repo *AdminService) CreateJobPosts(user *models.JobCreation) error {
	return repo.CreatePostForUser(user)
}

// update their job post by their IDs
func (repo *AdminService) UpdatePosts(user *models.JobCreation, jobid int) error {
	return repo.UpdateJobPost(user, jobid)
}

// get their postdetails jobDetailsBy role
func (repo *AdminService) GetAllPostsByRole(roleType string, roleID int) ([]models.UserJobDetails, error) {
	return repo.GetByRoles(roleType, roleID)
}

// get their applied details by their JobIds
func (repo *AdminService) GetAllPostsByJobID(jobid int, roleID int) ([]models.UserJobDetails, error) {
	return repo.GetByJobID(jobid, roleID)
}

// get their User's particular jobs By their userID's
func (repo *AdminService) GetAllPostsByUserID(userID int, roleID int) ([]models.UserJobDetails, error) {
	return repo.GetByUserID(userID, roleID)
}

// get thier own post details By admin
func (repo *AdminService) GetOwnPost(roleID int) ([]models.JobCreation, error) {
	return repo.GetAllOwnPosts(roleID)
}
