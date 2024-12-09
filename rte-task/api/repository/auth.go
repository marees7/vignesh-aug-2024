package repository

import (
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type AuthInterface interface {
	CheckEmailAddress(user *models.UserDetails, count int64) (int64, error)
	CreateDetailsByTheirRoles(user *models.UserDetails) error
	LoginEmailCheckExists(user *models.UserDetails, founduser *models.UserDetails) error
	CheckPhoneNumber(user *models.UserDetails, count int64) (int64, error)
}

type authrepo struct {
	*gorm.DB
}

// check email is exixts or not in DB
func (db *authrepo) CheckEmailAddress(user *models.UserDetails, count int64) (int64, error) {
	DbEmail := db.Model(&models.UserDetails{}).Where("email=?", user.Email).Count(&count)
	if DbEmail.Error != nil {
		return 0, DbEmail.Error
	}

	return count, nil
}

// create user details By their roles
func (db *authrepo) CreateDetailsByTheirRoles(user *models.UserDetails) error {
	dbvalues := db.Create(user)
	if dbvalues.Error != nil {
		return dbvalues.Error
	}

	return nil
}

// Check Email address while Login with their email ID
func (db *authrepo) LoginEmailCheckExists(user *models.UserDetails, founduser *models.UserDetails) error {
	value := db.Where("email=?", user.Email).First(&founduser)
	if value.Error != nil {
		return value.Error
	}

	return nil
}

// check phone number is exists or not in DB
func (databaseconnect *authrepo) CheckPhoneNumber(user *models.UserDetails, count int64) (int64, error) {
	DbPhone := databaseconnect.Model(&models.UserDetails{}).Where("phone_number=?", user.PhoneNumber).Count(&count)
	if DbPhone.Error != nil {
		return 0, DbPhone.Error
	}
	
	return count, nil
}
