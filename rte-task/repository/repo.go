package repository

import (
	"github.com/Vigneshwartt/golang-rte-task/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepsoitory(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (databaseconnect UserRepository) RepoEmailForm(user *models.UsersTable, count int64) (int64, error) {
	DbEmail := databaseconnect.DB.Model(&models.UsersTable{}).Where("email=?", user.Email).Count(&count)
	if DbEmail.Error != nil {
		return 0, DbEmail.Error
	}
	return count, nil
}

func (databaseconnect UserRepository) RepoPhoneForm(user *models.UsersTable, count int64) (int64, error) {
	DbPhone := databaseconnect.DB.Model(&models.UsersTable{}).Where("phone_number=?", user.PhoneNumber).Count(&count)
	if DbPhone.Error != nil {
		return 0, DbPhone.Error
	}
	return count, nil
}

func (databaseconnect UserRepository) RepoCreate(user *models.UsersTable) error {
	dbvalues := databaseconnect.DB.Create(user)
	if dbvalues.Error != nil {
		return dbvalues.Error
	}
	return nil
}

func (database UserRepository) RepoLoginEmail(user *models.UsersTable, founduser *models.UsersTable) error {
	value := database.DB.Where("email=?", user.Email).First(&founduser)
	if value.Error != nil {
		return value.Error
	}
	return nil
}

func (database UserRepository) RepoFindRoleID(user *models.UsersTable, founduser *models.UsersTable) error {
	value := database.DB.Where("role_id=?", founduser.RoleId).First(&founduser)
	if value.Error != nil {
		return value.Error
	}
	return nil
}

func (database UserRepository) RepoFindByUserID(user *models.UsersTable, founduser *models.UsersTable) error {
	value := database.DB.Where("role_id=?", founduser.RoleId).First(&founduser)
	if value.Error != nil {
		return value.Error
	}
	return nil
}

func (database UserRepository) RepoFindAllUsers(user *[]models.UsersTable) error {
	value := database.DB.Where("role_type=?", "USER").Find(&user)
	if value.Error != nil {
		return value.Error
	}
	return nil
}

func (database UserRepository) RepoFindSpecificID(user *models.UsersTable, roleid string) error {
	dbValue := database.DB.Where("role_id=?", roleid).Find(&user)
	if dbValue.Error != nil {
		return dbValue.Error
	}
	return nil
}

func (databaseconnect UserRepository) RepoCreateNewPost(user *models.JobCreation) error {
	dbvalues := databaseconnect.DB.Create(user)
	if dbvalues.Error != nil {
		return dbvalues.Error
	}
	return nil
}

func (databaseconnect UserRepository) RepoGetAllPosts(user *[]models.JobCreation) error {
	dbvalues := databaseconnect.DB.Find(&user)
	if dbvalues.Error != nil {
		return dbvalues.Error
	}
	return nil
}

func (database UserRepository) RepoGetByJobRole(user *[]models.JobCreation, jobs string) error {
	dbValue := database.DB.Where(&models.JobCreation{JobTitle: jobs}).Find(&user)
	if dbValue.Error != nil {
		return dbValue.Error
	}
	return nil
}

func (database UserRepository) RepoGetByCountryDetails(user *[]models.JobCreation, country string) error {
	dbvalue := database.DB.Where(&models.JobCreation{Country: country}).Find(&user)
	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	return nil
}

func (databaseconnect UserRepository) RepoApplyJobPost(user *models.UserJobDetails) error {
	dbvalues := databaseconnect.DB.Create(user)
	if dbvalues.Error != nil {
		return dbvalues.Error
	}
	return nil
}

func (databaseconnectdata UserRepository) RepoGetJobAppliedDetails(user *[]models.UserJobDetails, roletype string) error {
	dbvalue := databaseconnectdata.DB.Where(&models.UserJobDetails{JobRole: roletype}).Find(&user)
	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	return nil
}

func (databaseconnectdata UserRepository) RepoGetJobAppliedAllJobs(user *[]models.UserJobDetails, roletype string) error {
	dbvalue := databaseconnectdata.DB.Where(&models.UserJobDetails{JobRole: roletype}).Find(&user)
	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	return nil
}
