package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

//	func (database UserRepository) RepoFindByUserID(user *models.UsersTable, founduser *models.UsersTable) error {
//		value := database.DB.Where("role_id=?", founduser.RoleId).First(&founduser)
//		if value.Error != nil {
//			return value.Error
//		}
//		return nil
//	}
type UserRepo interface {
	RepoGetAllPosts(user *[]models.JobCreation) error
	RepoGetByJobRole(user *[]models.JobCreation, jobs string) error
	RepoGetByCountryDetails(user *[]models.JobCreation, country string) error
	RepoApplyJobPost(user *models.UserJobDetails, roletype string, userid int, applyuserid int) error
	GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int) error
	UserGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int) error
}

type userrepo struct {
	*gorm.DB
}

func (db *userrepo) RepoGetAllPosts(user *[]models.JobCreation) error {
	dbvalues := db.Find(&user)
	if dbvalues.Error != nil {
		return dbvalues.Error
	}
	return nil
}

func (db *userrepo) RepoGetByJobRole(user *[]models.JobCreation, jobs string) error {
	data := db.Where("job_title=?", jobs).First(&user)
	if data.Error != nil {
		return data.Error
	}

	dbValue := db.Where(&models.JobCreation{JobTitle: jobs}).Find(&user)
	if dbValue.Error != nil {
		return dbValue.Error
	}
	return nil
}

func (db *userrepo) RepoGetByCountryDetails(user *[]models.JobCreation, country string) error {
	dbvalue := db.Where(&models.JobCreation{Country: country}).Find(&user)
	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	return nil
}

func (db *userrepo) RepoApplyJobPost(user *models.UserJobDetails, roletype string, userid int, applyuserid int) error {
	if roletype == "USER" && applyuserid == userid {
		dbvalues := db.Create(user)
		if dbvalues.Error != nil {
			return dbvalues.Error
		} else if dbvalues.RowsAffected == 0 {
			return fmt.Errorf("could not apply post id:%d", userid)
		}
	} else {
		return fmt.Errorf("could not able to apply post")
	}
	return nil
}

func (db *userrepo) GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int) error {
	data := db.Where("user_id=?", userId).First(&user)
	if data.Error != nil {
		return data.Error
	}

	dbvalue := db.
		Preload("Job").
		Where("user_id = ?", userId).
		Find(&user)

	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	return nil
}

func (db *userrepo) UserGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int) error {
	// var newuser models.UserJobDetails
	data := db.Where("user_id=?", userId).First(&user)
	if data.Error != nil {
		return data.Error
	}
	// if roletype == "USER" && newuser.UserID == userId {
	dbvalue := db.
		Preload("Job").
		Where("user_id = ?", userId).
		Find(&user)

	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	// }
	return nil
}
