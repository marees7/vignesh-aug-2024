package repository

import "gorm.io/gorm"

type UserRepository struct {
	Admin AdminRepository
	Auth  AuthInterface
	User  UserRepo
}

func GetAdminRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Admin: &adminRepo{db},
	}
}

func GetAuthRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		Auth: &authrepo{db},
	}
}
func GetUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		User: &userrepo{db},
	}
}
