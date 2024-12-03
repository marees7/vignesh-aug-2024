package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"gorm.io/gorm"
)

type AdminRepository interface {
	RepoFindAllUsers(user *[]models.UserDetails) error
	RepoCreateNewPost(user *models.JobCreation, roletype string, userid int, applyuserid int) error
	RepoUpdateJobPost(user *models.JobCreation, jobID int, role string, userid int, useridvalues int) error
	RepoGetJobAppliedDetailsbyrole(user *[]models.UserJobDetails, roletype string, usertype string, userid int, applyuserid int) error
	// RepoGetJobAppliedAllJobs(user *[]models.UserJobDetails, roletype string, userid int, applyuserid int)
	RepoGetJobAppliedDetailsByJobId(user *[]models.UserJobDetails, roletype string, userid int, jobID int, applyuserid int) error
	RepoGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int, usertype string, applyid int, adminid int) error
}
type adminRepo struct {
	*gorm.DB
}

func (database *adminRepo) RepoFindAllUsers(user *[]models.UserDetails) error {
	value := database.Where("role_type=?", "USER").Find(&user)
	if value.Error != nil {
		return value.Error
	}
	return nil
}

func (database *adminRepo) RepoCreateNewPost(user *models.JobCreation, roletype string, userid int, applyuserid int) error {
	if roletype == "ADMIN" && applyuserid == userid {
		dbvalues := database.Create(user)
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

func (database *adminRepo) RepoUpdateJobPost(user *models.JobCreation, jobID int, role string, userid int, useridvalues int) error {
	var existingPost models.JobCreation
	result := database.Where("job_id= ?", jobID).First(&existingPost)
	if result.Error != nil {
		return result.Error
	}

	if role == "ADMIN" && userid == useridvalues {
		updateResult := database.Model(&models.JobCreation{}).Where("job_id= ?", jobID).Updates(map[string]interface{}{
			"JobStatus": user.JobStatus,
			"JobTime":   user.JobTime,
			"Vacancy":   user.Vacancy,
		})

		if updateResult.Error != nil {
			return updateResult.Error
		} else if updateResult.RowsAffected == 0 {
			return fmt.Errorf("no rows affected for job_id: %d", jobID)
		}
	} else {
		return fmt.Errorf("unauthorized role: %s", role)
	}

	return nil
}

func (database *adminRepo) RepoGetJobAppliedDetailsbyrole(user *[]models.UserJobDetails, roletype string, usertype string, userid int, applyuserid int) error {
	data := database.Where("job_role=?", roletype).First(&user)
	if data.Error != nil {
		return data.Error
	}
	if usertype == "ADMIN" && userid == applyuserid {
		dbvalue := database.Preload("User").
			Where(&models.UserJobDetails{JobRole: roletype}).Find(&user)
		if dbvalue.Error != nil {
			return dbvalue.Error
		}
	} else {
		return fmt.Errorf("unauthorized role: %d", userid)
	}
	return nil
}

// func (database *adminRepo) RepoGetJobAppliedAllJobs(user *[]models.UserJobDetails, roletype string, userid int, applyuserid int) error {
// 	if roletype == "ADMIN" && applyuserid == userid {
// 		dbvalue := database.Preload("User").Find(&user)
// 		if dbvalue.Error != nil {
// 			return dbvalue.Error
// 		}
// 	} else {
// 		return fmt.Errorf("could not able to apply post")
// 	}
// 	return nil
// }

func (database *adminRepo) RepoGetJobAppliedDetailsByJobId(user *[]models.UserJobDetails, roletype string, userid int, jobID int, applyuserid int) error {
	data := database.Where("job_id=?", jobID).First(&user)
	if data.Error != nil {
		return data.Error
	}

	if roletype == "ADMIN" && applyuserid == userid {
		dbvalue := database.
			Preload("User").
			Where("job_id = ?", jobID).
			Find(&user)

		if dbvalue.Error != nil {
			return dbvalue.Error
		}
	} else {
		return fmt.Errorf("could get that post")
	}
	return nil

}

func (database *adminRepo) RepoGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int, usertype string, applyid int, adminid int) error {
	data := database.Where("user_id=?", userId).First(&user)
	if data.Error != nil {
		return data.Error
	}
	if usertype == "ADMIN" && applyid == adminid {
		dbvalue := database.
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

// func (databaseconnectdata UserRepository) RepoDeleteJobPost(user *models.JobCreation, jobID int, role string) error {
// 	var existingPost models.JobCreation
// 	result := databaseconnectdata.DB.Where("job_id = ?", jobID).First(&existingPost)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if role == "ADMIN" {
// 		updateResult := databaseconnectdata.DB.Model(&models.JobCreation{}).Where("job_id = ?", jobID).Delete(&user)
// 		if updateResult.Error != nil {
// 			return updateResult.Error
// 		} else if updateResult.RowsAffected == 0 {
// 			return fmt.Errorf("no rows affected for job_id: %d", jobID)
// 		}
// 	} else {
// 		return fmt.Errorf("unauthorized role: %s", role)
// 	}
// 	return nil
// }
