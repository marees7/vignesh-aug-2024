package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	RepoGetAllPosts(user *[]models.JobCreation, usertype string) error
	RepoGetByJobRole(user *[]models.JobCreation, jobs string, country string, usertype string) error
	RepoApplyJobPost(user *models.UserJobDetails, roletype string, userid int, applyuserid int) error
	UserGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int, tokenid int) error
}

type userrepo struct {
	*gorm.DB
}

func (db *userrepo) RepoGetAllPosts(user *[]models.JobCreation, usertype string) error {
	if usertype == "ADMIN" || usertype == "USER" {
		dbvalues := db.Find(&user)
		if dbvalues.Error != nil {
			return dbvalues.Error
		}
	} else {
		return fmt.Errorf("could not able to Get the details")
	}
	return nil
}

func (db *userrepo) RepoGetByJobRole(user *[]models.JobCreation, jobs string, country string, usertype string) error {
	if usertype == "ADMIN" || usertype == "USER" {
		data := db.Where("job_title=?", jobs).First(&user)
		if data.Error != nil {
			return data.Error
		}
		dbvalue := db.Where(&models.JobCreation{Country: country}).First(&user)
		if dbvalue.Error != nil {
			return dbvalue.Error
		}
		dbValue := db.Where(&models.JobCreation{JobTitle: jobs, Country: country}).Find(&user)
		if dbValue.Error != nil {
			return dbValue.Error
		}
		return nil
	} else {
		return fmt.Errorf("could not able to Get the details")
	}

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
		return fmt.Errorf("could not able to get details")
	}
	return nil
}

func (db *userrepo) UserGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int, tokenid int) error {
	data := db.Where("user_id=?", userId).First(&user)
	if data.Error != nil {
		return data.Error
	}
	if tokenid == userId {
		dbvalue := db.
			Preload("Job").
			Where("user_id = ?", userId).
			Find(&user)

		if dbvalue.Error != nil {
			return dbvalue.Error
		}
	} else {
		return fmt.Errorf("could get that post")
	}
	return nil
}
