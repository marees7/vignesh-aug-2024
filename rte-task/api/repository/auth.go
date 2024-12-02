package repository

import (
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type AuthInterface interface {
	RepoEmailForm(user *models.UserDetails, count int64) (int64, error)
	RepoCreate(user *models.UserDetails) error
	RepoLoginEmail(user *models.UserDetails, founduser *models.UserDetails) error
	RepoFindRoleID(user *models.UserDetails, founduser *models.UserDetails) error
}

type authrepo struct {
	*gorm.DB
}

func (db *authrepo) RepoEmailForm(user *models.UserDetails, count int64) (int64, error) {
	DbEmail := db.Model(&models.UserDetails{}).Where("email=?", user.Email).Count(&count)
	if DbEmail.Error != nil {
		return 0, DbEmail.Error
	}
	return count, nil
}


func (db *authrepo) RepoCreate(user *models.UserDetails) error {
	dbvalues := db.Create(user)
	if dbvalues.Error != nil {
		return dbvalues.Error
	}
	return nil
}

func (db *authrepo) RepoLoginEmail(user *models.UserDetails, founduser *models.UserDetails) error {
	value := db.Where("email=?", user.Email).First(&founduser)
	if value.Error != nil {
		return value.Error
	}
	return nil
}

func (db *authrepo) RepoFindRoleID(user *models.UserDetails, founduser *models.UserDetails) error {
	value := db.Where("role_id=?", founduser.RoleId).First(&founduser)
	if value.Error != nil {
		return value.Error
	}
	return nil
}
