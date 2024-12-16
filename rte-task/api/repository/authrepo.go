package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/internals"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
)

type I_AuthRepo interface {
	GetValidEmailAddress(user *models.UserDetails) error
	CreateUsers(user *models.UserDetails) error
	GetUserMails(user *models.UserDetails, founduser *models.UserDetails) error
	GetPhoneNumber(user *models.UserDetails) error
}

type AuthRepo struct {
	*internals.ConnectionNew
}

func GetAuthRepository(db *internals.ConnectionNew) I_AuthRepo {
	return &AuthRepo{
		db,
	}
}

// check email is exixts or not in DB
func (database *AuthRepo) GetValidEmailAddress(user *models.UserDetails) error {
	var count int64

	db := database.Model(&models.UserDetails{}).Where("email=?", user.Email).Count(&count)
	if db.Error != nil {
		return fmt.Errorf("error occured while searching the email ")
	}
	if count > 0 {
		return fmt.Errorf(" error Email is already registered by someones's ,Try another mail")
	}
	return nil
}

// create user details By their roles
func (database *AuthRepo) CreateUsers(user *models.UserDetails) error {
	db := database.Create(user)
	if db.Error != nil {
		return fmt.Errorf("can't able to create your data")
	}
	return nil
}

// Check Email address while Login with their email ID
func (database *AuthRepo) GetUserMails(user *models.UserDetails, founduser *models.UserDetails) error {
	db := database.Where("email=?", user.Email).First(&founduser)
	if db.Error != nil {
		return fmt.Errorf("cant' Match your MailId properly,check it once")
	}
	return nil
}

// check phone number is exists or not in DB
func (database *AuthRepo) GetPhoneNumber(user *models.UserDetails) error {
	var count int64

	db := database.Model(&models.UserDetails{}).Where("phone_number=?", user.PhoneNumber).Count(&count)
	if db.Error != nil {
		return fmt.Errorf("cant' able to create your Phone number")
	}
	if count > 0 {
		return fmt.Errorf(" error Phone Number is already registered by someones's ,Please verify it")
	}
	return nil
}
