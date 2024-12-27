package service

import (
	"testing"

	"github.com/Vigneshwartt/golang-rte-task/common/dto"
	"github.com/Vigneshwartt/golang-rte-task/mocks"
	"github.com/Vigneshwartt/golang-rte-task/pkg/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MockAuthRepoTestSuite struct {
	suite.Suite
	mockRepo      *mocks.IAuthRepo
	signUpService IAuthService
}

func (suite *MockAuthRepoTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.IAuthRepo)
	suite.signUpService = InitAuthService(suite.mockRepo)
}

func (suite *MockAuthRepoTestSuite) TestSignUp() {
	userDetails := models.UserDetails{
		Name:        "vikki",
		Email:       "a@gmail.com",
		Password:    "Abcd123&",
		PhoneNumber: "1234567890",
		RoleType:    "USER",
	}
	err := &dto.ErrorResponse{}
	data := models.UserDetails{
		Name:        "vikki",
		Email:       "a@gmail.com",
		Password:    "Abcd123&",
		PhoneNumber: "1234567890",
		RoleType:    "USER",
	}
	suite.mockRepo.On("CreateUser", &data).Return(err)
	errors := suite.signUpService.CreateUser(&userDetails)

	suite.mockRepo.AssertExpectations(suite.T())
	assert.Equal(suite.T(), userDetails, data)
	assert.Equal(suite.T(), err, errors)

}

func (suite *MockAuthRepoTestSuite) TestLoginUp() {
	userLoginDetails := &models.UserDetails{
		Name:        "vikki",
		Email:       "a@gmail.com",
		Password:    "Abcd123&",
		PhoneNumber: "1234567890",
		RoleType:    "USER",
	}
	err := &dto.ErrorResponse{}

	suite.mockRepo.On("GetUserDetail", userLoginDetails).Return(&models.UserDetails{
		Name:        "vikki",
		Email:       "a@gmail.com",
		Password:    "Abcd123&",
		PhoneNumber: "1234567890",
		RoleType:    "USER",
	}, err)

	value, errs := suite.signUpService.GetUserDetail(userLoginDetails)

	suite.mockRepo.AssertExpectations(suite.T())

	assert.NotNil(suite.T(), errs)
	assert.Equal(suite.T(), userLoginDetails, value)
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(MockAuthRepoTestSuite))
}
