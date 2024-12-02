package repository

import (
	"fmt"

	"github.com/Vigneshwartt/golang-rte-task/models"
)

func (databaseconnect UserRepository) RepoCreateNewPost(user *models.JobCreation) error {
	dbvalues := databaseconnect.DB.Create(user)
	if dbvalues.Error != nil {
		return dbvalues.Error
	}
	return nil
}

func (databaseconnectdata UserRepository) RepoUpdateJobPost(user *models.JobCreation, jobID int, role string) error {
	var existingPost models.JobCreation
	result := databaseconnectdata.DB.Where("job_id_new = ?", jobID).First(&existingPost)
	if result.Error != nil {
		return result.Error
	}

	if role == "ADMIN" {
		updateResult := databaseconnectdata.DB.Model(&models.JobCreation{}).Where("job_id_new = ?", jobID).Updates(map[string]interface{}{
			"JobStatus": user.JobStatus,
			"JobTime":   user.JobTime,
			"Vacancy":   user.Vacancy,
		})

		if updateResult.Error != nil {
			return updateResult.Error
		} else if updateResult.RowsAffected == 0 {
			return fmt.Errorf("no rows affected for job_id_new: %d", jobID)
		}
	} else {
		return fmt.Errorf("unauthorized role: %s", role)
	}

	return nil
}

func (databaseconnectdata UserRepository) RepoGetJobAppliedDetailsbyrole(user *[]models.UserJobDetails, roletype string) error {
	data := databaseconnectdata.DB.Where("job_role=?", roletype).First(&user)
	if data.Error != nil {
		return data.Error
	}
	dbvalue := databaseconnectdata.DB.Preload("User").
		Where(&models.UserJobDetails{JobRole: roletype}).Find(&user)
	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	return nil
}

func (databaseconnectdata UserRepository) RepoGetJobAppliedAllJobs(user *[]models.UserJobDetails) error {
	dbvalue := databaseconnectdata.DB.Preload("User").Find(&user)
	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	return nil
}

func (databaseconnectdata UserRepository) RepoGetJobAppliedDetailsByJobId(user *[]models.UserJobDetails, jobID int) error {
	data := databaseconnectdata.DB.Where("job_id=?", jobID).First(&user)
	if data.Error != nil {
		return data.Error
	}

	dbvalue := databaseconnectdata.DB.
		Preload("User").
		Where("job_id = ?", jobID).
		Find(&user)

	if dbvalue.Error != nil {
		return dbvalue.Error
	}
	return nil
}

func (databaseconnectdata UserRepository) RepoGetJobAppliedDetailsByUserId(user *[]models.UserJobDetails, userId int) error {
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
