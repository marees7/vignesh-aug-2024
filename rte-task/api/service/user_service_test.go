package service

import (
	"testing"

	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/mocks"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockUserTestSuite struct {
	suite.Suite
	mockRepo    *mocks.IUserRepo
	userService IUserService
}

func (suite *MockUserTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.IUserRepo)
	suite.userService = InitUserService(suite.mockRepo)
}

func (suite *MockUserTestSuite) TestCreateApplicant() {
	// JobId:=&1
	userApplicantPost := models.UserJobDetails{
		UserID:     1,
		Experience: 2,
		Skills:     "Golang",
		Language:   "English",
		Country:    "India",
		JobRole:    "HR",
	}
	err := &dto.ErrorResponse{}

	suite.mockRepo.On("CreateApplication", &userApplicantPost).Return(err)
	errors := suite.userService.CreateApplication(&userApplicantPost)
	suite.mockRepo.AssertExpectations(suite.T())
	assert.Equal(suite.T(), err, errors)
}

func (suite *MockUserTestSuite) TestGetAllJobPosts() {
	var count int64

	jobCreateDetails := []models.JobCreation{
		{JobID: 1, AdminID: 1, CompanyName: "ABC", CompanyEmail: "abc@gmail.com", JobRole: "HR", JobStatus: "ON GOING", JobTime: "PART TIME", Experience: "0-2 years", Description: "young dynamic persons need for this work",
			Skills: "Go", Vacancy: 20, Country: "India", Address: models.Address{Street: "abc street", City: "India", State: "TamilNadu", ZipCode: "606106"}},
	}

	err := &dto.ErrorResponse{}

	suite.mockRepo.On("GetAllJobPosts", map[string]interface{}{}).Return([]models.JobCreation{
		{JobID: 1, AdminID: 1, CompanyName: "ABC", CompanyEmail: "abc@gmail.com", JobRole: "HR", JobStatus: "ON GOING", JobTime: "PART TIME", Experience: "0-2 years", Description: "young dynamic persons need for this work",
			Skills: "Go", Vacancy: 20, Country: "India", Address: models.Address{Street: "abc street", City: "India", State: "TamilNadu", ZipCode: "606106"}},
	}, err, count)

	value, errors, _ := suite.userService.GetAllJobPosts(map[string]interface{}{})

	suite.mockRepo.AssertExpectations(suite.T())

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors, err)
	assert.Equal(suite.T(), jobCreateDetails, value)

}

func (suite *MockUserTestSuite) TestGetUserAppliedJobs() {
	var count int64
	err := &dto.ErrorResponse{}

	userJobs := []models.UserJobDetails{
		{UserID: 1,
			Experience: 2,
			Skills:     "Go",
			Language:   "English",
			Country:    "India",
			JobRole:    "HR"}}

	suite.mockRepo.On("GetUserAppliedJobs", map[string]interface{}{}).Return([]models.UserJobDetails{
		{UserID: 1,
			Experience: 2,
			Skills:     "Go",
			Language:   "English",
			Country:    "India",
			JobRole:    "HR"},
	}, err, count)

	value, errors, _ := suite.userService.GetUserAppliedJobs(map[string]interface{}{})

	suite.mockRepo.AssertExpectations(suite.T())

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), errors, err)
	assert.Equal(suite.T(), userJobs, value)
}

func TestUserTestSuites(t *testing.T) {
	suite.Run(t, new(MockUserTestSuite))
}
