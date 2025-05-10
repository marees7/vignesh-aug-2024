package service

import (
	"testing"

	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/mocks"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockAdminTestSuite struct {
	suite.Suite
	mockRepo     *mocks.IAdminRepo
	adminService IAdminService
}

func (suite *MockAdminTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.IAdminRepo)
	suite.adminService = InitAdminService(suite.mockRepo)
}

func (suite *MockAdminTestSuite) TestCreateJobPost() {
	userJobPost := models.JobCreation{
		AdminID:      1,
		CompanyName:  "ABC",
		CompanyEmail: "abc@gmail.com",
		JobRole:      "Java",
		JobStatus:    "ON GOING",
		JobTime:      "PART TIME",
		Description:  "young dynamic persons need for this work",
		Experience:   "0-2 years",
		Skills:       "Go,Java",
		Vacancy:      3,
		Country:      "India",
		Address: models.Address{
			Street:  "abc street",
			City:    "madurai",
			State:   "TamilNadu",
			ZipCode: "123456",
		},
	}
	err := &dto.ErrorResponse{}
	data := models.JobCreation{
		AdminID:      1,
		CompanyName:  "ABC",
		CompanyEmail: "abc@gmail.com",
		JobRole:      "Java",
		JobStatus:    "ON GOING",
		JobTime:      "PART TIME",
		Description:  "young dynamic persons need for this work",
		Experience:   "0-2 years",
		Skills:       "Go,Java",
		Vacancy:      3,
		Country:      "India",
		Address: models.Address{
			Street:  "abc street",
			City:    "madurai",
			State:   "TamilNadu",
			ZipCode: "123456",
		},
	}
	suite.mockRepo.On("CreateJobPost", &userJobPost).Return(err)
	errors := suite.adminService.CreateJobPost(&userJobPost)

	suite.mockRepo.AssertExpectations(suite.T())
	assert.Equal(suite.T(), userJobPost, data)
	assert.Equal(suite.T(), err, errors)
}

func (suite *MockAdminTestSuite) TestGetJobsAppliedByUser() {
	var count int64
	err := &dto.ErrorResponse{}

	userjobDetail := []models.UserJobDetails{
		{UserID: 1,
			Experience: 2,
			Skills:     "Go",
			Language:   "English",
			Country:    "India",
			JobRole:    "HR"},

		{UserID: 1,
			Experience: 1,
			Skills:     "Java",
			Language:   "English",
			Country:    "USA",
			JobRole:    "Admin"},
	}

	suite.mockRepo.On("GetJobsAppliedByUser", map[string]interface{}{}).Return([]models.UserJobDetails{
		{UserID: 1,
			Experience: 2,
			Skills:     "Go",
			Language:   "English",
			Country:    "India",
			JobRole:    "HR"},

		{UserID: 1,
			Experience: 1,
			Skills:     "Java",
			Language:   "English",
			Country:    "USA",
			JobRole:    "Admin"},
	}, err, count)

	value, errors, _ := suite.adminService.GetJobsAppliedByUser(map[string]interface{}{})

	suite.mockRepo.AssertExpectations(suite.T())

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors, err)
	assert.Equal(suite.T(), userjobDetail, value)
}

func (suite *MockAdminTestSuite) TestUpdateJobPost() {
	var jobId, userId int
	userUpdatePost := models.JobCreation{
		JobID:     1,
		JobStatus: "COMPLETED",
		Vacancy:   0,
	}
	err := &dto.ErrorResponse{}
	data := models.JobCreation{
		JobID:     1,
		JobStatus: "COMPLETED",
		Vacancy:   0,
	}
	suite.mockRepo.On("UpdateJobPost", &userUpdatePost, jobId, userId).Return(err)
	errors := suite.adminService.UpdateJobPost(&userUpdatePost, jobId, userId)

	suite.mockRepo.AssertExpectations(suite.T())
	assert.Equal(suite.T(), userUpdatePost, data)
	assert.Equal(suite.T(), err, errors)
}

func (suite *MockAdminTestSuite) TestJobsCreated() {
	var count int64
	userJobPost := []models.JobCreation{
		{
			AdminID:      1,
			CompanyName:  "ABC",
			CompanyEmail: "abc@gmail.com",
			JobRole:      "Java",
			JobStatus:    "ON GOING",
			JobTime:      "PART TIME",
			Description:  "young dynamic persons need for this work",
			Experience:   "0-2 years",
			Skills:       "Go,Java",
			Vacancy:      3,
			Country:      "India",
			Address: models.Address{
				Street:  "abc street",
				City:    "madurai",
				State:   "TamilNadu",
				ZipCode: "123456",
			}},
	}
	err := &dto.ErrorResponse{}

	suite.mockRepo.On("GetJobsCreated", map[string]interface{}{}).Return([]models.JobCreation{{
		AdminID:      1,
		CompanyName:  "ABC",
		CompanyEmail: "abc@gmail.com",
		JobRole:      "Java",
		JobStatus:    "ON GOING",
		JobTime:      "PART TIME",
		Description:  "young dynamic persons need for this work",
		Experience:   "0-2 years",
		Skills:       "Go,Java",
		Vacancy:      3,
		Country:      "India",
		Address: models.Address{
			Street:  "abc street",
			City:    "madurai",
			State:   "TamilNadu",
			ZipCode: "123456",
		}},
	}, err, count)

	data, errors, _ := suite.adminService.GetJobsCreated(map[string]interface{}{})

	suite.mockRepo.AssertExpectations(suite.T())

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors, err)
	assert.Equal(suite.T(), userJobPost, data)
}
func (suite *MockAdminTestSuite) TestDeleteJobPost() {
	err := &dto.ErrorResponse{}

	suite.mockRepo.On("DeleteJobPost").Return(err)

	errors := suite.adminService.DeleteJobPost()

	suite.mockRepo.AssertExpectations(suite.T())
	assert.Equal(suite.T(), err, errors)
}

func TestAdminTestSuites(t *testing.T) {
	suite.Run(t, new(MockAdminTestSuite))
}
