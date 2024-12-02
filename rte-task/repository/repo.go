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

// func (database UserRepository) RepoFindByUserID(user *models.UsersTable, founduser *models.UsersTable) error {
// 	value := database.DB.Where("role_id=?", founduser.RoleId).First(&founduser)
// 	if value.Error != nil {
// 		return value.Error
// 	}
// 	return nil
// }

func (database UserRepository) RepoFindAllUsers(user *[]models.UsersTable) error {
	value := database.DB.Where("role_type=?", "USER").Find(&user)
	if value.Error != nil {
		return value.Error
	}
	return nil
}

func (database UserRepository) RepoFindSpecificID(user *models.UsersTable, roleid string) error {
	data := database.DB.Where("role_id=?", roleid).First(&user)
	if data.Error != nil {
		return data.Error
	}
	dbValue := database.DB.Where("role_id=?", roleid).Find(&user)
	if dbValue.Error != nil {
		return dbValue.Error
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
	data := database.DB.Where("job_title=?", jobs).First(&user)
	if data.Error != nil {
		return data.Error
	}

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

func (databaseconnect UserRepository) RepoApplyJobPost(user *models.UserJobDetails, roletype string) error {
	if roletype == "USER" {
		dbvalues := databaseconnect.DB.Create(user)
		if dbvalues.Error != nil {
			return dbvalues.Error
		}
	}
	return nil
}

func (databaseconnectdata UserRepository) GetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int) error {
	data := databaseconnectdata.DB.Where("user_id=?", userId).First(&user)
	if data.Error != nil {
		return data.Error
	}

	dbvalue := databaseconnectdata.DB.
		Preload("Job").
		Where("user_id = ?", userId).
		Find(&user)

	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	return nil
}

// func (databaseconnectdata UserRepository) RepoDeleteJobPost(user *models.JobCreation, jobID int, role string) error {
// 	var existingPost models.JobCreation
// 	result := databaseconnectdata.DB.Where("job_id_new = ?", jobID).First(&existingPost)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	if role == "ADMIN" {
// 		updateResult := databaseconnectdata.DB.Model(&models.JobCreation{}).Where("job_id_new = ?", jobID).Delete(&user)
// 		if updateResult.Error != nil {
// 			return updateResult.Error
// 		} else if updateResult.RowsAffected == 0 {
// 			return fmt.Errorf("no rows affected for job_id_new: %d", jobID)
// 		}
// 	} else {
// 		return fmt.Errorf("unauthorized role: %s", role)
// 	}
// 	return nil
// }
