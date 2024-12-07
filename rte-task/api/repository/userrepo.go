package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	RepoGetAllPosts(user *[]models.JobCreation, usertype string) error
	RepoGetByJobRole(user *[]models.JobCreation, jobs string, country string, usertype string) error
	RepoGetByCompanyName(user *[]models.JobCreation, company string, usertype string) error
	RepoApplyJobPost(user *models.UserJobDetails) error
	UserGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int, tokenid int) error
	CheckUserJobId(user *models.UserJobDetails, newpost *models.JobCreation) error
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
			return fmt.Errorf("cant'able to find your jobs Properly,Give him correctly")
		}
		dbvalue := db.Where(&models.JobCreation{Country: country}).First(&user)
		if dbvalue.Error != nil {
			return fmt.Errorf("cant'able to find your country Properly,Give him correctly")
		}
		dbValue := db.Where(&models.JobCreation{JobTitle: jobs, Country: country}).Find(&user)
		if dbValue.Error != nil {
			return fmt.Errorf("cant'able to find your details properly,Give him correctly")
		}
	} else {
		return fmt.Errorf("could not able to Get the details")
	}
	return nil
}

func (db *userrepo) RepoGetByCompanyName(user *[]models.JobCreation, company string, usertype string) error {
	if usertype == "ADMIN" || usertype == "USER" {
		data := db.Where("company_name=?", company).First(&user)
		if data.Error != nil {
			return fmt.Errorf("cant'able to find your jobs in that company,Give him correctly")
		}

		dbValue := db.Where(&models.JobCreation{CompanyName: company}).Find(&user)
		if dbValue.Error != nil {
			return fmt.Errorf("cant'able to find your jobs in that company,Give him correctly")
		}
	} else {
		return fmt.Errorf("could not able to Get the details")
	}
	return nil
}

func (db *userrepo) RepoApplyJobPost(user *models.UserJobDetails) error {
	dbvalues := db.Create(user)
	if dbvalues.Error != nil {
		return fmt.Errorf("can't able to apply the job post here,please Check")
	}
	return nil
}

func (db *userrepo) UserGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int, tokenid int) error {
	data := db.Where("user_id=?", userId).First(&user)
	if data.Error != nil {
		return fmt.Errorf("cant'able to find your UserId ,Give him correctly")
	}
	if tokenid == userId {
		dbvalue := db.
			Preload("Job").
			Where("user_id = ?", userId).
			Find(&user)

		if dbvalue.Error != nil {
			return fmt.Errorf("cant'able to create your details ,Give him correctly")
		}
	} else {
		return fmt.Errorf("your payload Id and user Id is mismatching here,check it once")
	}
	return nil
}

func (db *userrepo) CheckUserJobId(user *models.UserJobDetails, newuser *models.JobCreation) error {
	var count int64

	jobid := db.Model(&models.JobCreation{}).Where("job_id=? ", user.JobID).First(&newuser)
	if jobid.Error != nil {
		return fmt.Errorf("unable to fetch Job Details properly,Check it JobId once")
	}
	if newuser.JobStatus == "COMPLETED" {
		return fmt.Errorf("this Job Application is closed")
	}

	data := db.Model(&models.UserJobDetails{}).Where("user_id=? AND job_id=?", user.UserId, user.JobID).Count(&count)
	if count > 0 {
		return fmt.Errorf("already registered,You have applied for this job")
	}
	if data.Error != nil {
		return fmt.Errorf("error occured,While applying the post")
	}
	return nil
}
