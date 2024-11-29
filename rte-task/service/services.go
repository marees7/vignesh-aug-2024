package service

import (
	"github.com/Vigneshwartt/golang-rte-task/models"
	"github.com/Vigneshwartt/golang-rte-task/repository"
)

type UserService struct {
	DB repository.UserRepository
}

func NewUserService(db repository.UserRepository) UserService {
	return UserService{DB: db}
}

func (service UserService) ServiceRepoemail(user *models.UsersTable, count int64) (int64, error) {
	return service.DB.RepoEmailForm(user, count)
}

func (service UserService) ServicePhoneForm(user *models.UsersTable, count int64) (int64, error) {
	return service.DB.RepoPhoneForm(user, count)
}

func (service UserService) ServiceCreate(user *models.UsersTable) error {
	return service.DB.RepoCreate(user)
}

func (service UserService) ServiceLoginEmail(user *models.UsersTable, founduser *models.UsersTable) error {
	return service.DB.RepoLoginEmail(user, founduser)
}

func (service UserService) ServiceFindRoleID(user *models.UsersTable, founduser *models.UsersTable) error {
	return service.DB.RepoFindRoleID(user, founduser)
}

func (service UserService) ServiceFindAllUsers(user *[]models.UsersTable) error {
	return service.DB.RepoFindAllUsers(user)
}

func (service UserService) ServiceFindSpecificUser(user *models.UsersTable, roleid string) error {
	return service.DB.RepoFindSpecificID(user, roleid)
}

func (service UserService) ServiceCreatePost(user *models.JobCreation) error {
	return service.DB.RepoCreateNewPost(user)
}

func (service UserService) ServiceGetAllPostDetails(user *[]models.JobCreation) error {
	return service.DB.RepoGetAllPosts(user)
}

func (service UserService) ServiceGetJobDetailsByRole(user *[]models.JobCreation, jobrole string) error {
	return service.DB.RepoGetByJobRole(user, jobrole)
}

func (service UserService) ServiceGetDetailsByCountry(user *[]models.JobCreation, country string) error {
	return service.DB.RepoGetByCountryDetails(user, country)
}

func (service UserService) ApplyJobPost(user *models.UserJobDetails) error {
	
	return service.DB.RepoApplyJobPost(user)
}

func (service UserService) ServiceGetJobAppliedDetails(user *[]models.UserJobDetails, roletype string) error {
	return service.DB.RepoGetJobAppliedDetails(user, roletype)
}

func (service UserService) ServiceGetJobAppliedAllJobs(user *[]models.UserJobDetails, roletype string) error {
	return service.DB.RepoGetJobAppliedAllJobs(user, roletype)
}
