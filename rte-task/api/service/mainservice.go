package service

import "github.com/Vigneshwartt/golang-rte-task/api/repository"

type UserService struct {
	User  UserServices
	Admin AdminService
	Auth  AuthService
}

func GetAdminService(db *repository.UserRepository) *UserService {
	return &UserService{
		Admin: &adminservice{db},
	}
}
func GetUserService(db *repository.UserRepository) *UserService {
	return &UserService{
		User: &userservice{db},
	}
}
func GetAuthService(db *repository.UserRepository) *UserService {
	return &UserService{
		Auth: &authservice{db},
	}
}
