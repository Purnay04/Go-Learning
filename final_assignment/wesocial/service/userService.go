package service

import (
	"fmt"
	"wesocial/repo"
)

type UserService interface {
	RegisterUser(*repo.User) (*repo.User, error)
	LoginUser(string, string) (*repo.User, error)
}

type UserServiceImpl struct {
	userRepo repo.UserRepository
}

func NewUserService(userRepo repo.UserRepository) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (userService *UserServiceImpl) RegisterUser(newUser *repo.User) (*repo.User, error) {
	if newUser.Name == "" {
		return nil, fmt.Errorf("user name is not provided")
	}
	if newUser.Email == "" {
		return nil, fmt.Errorf("user email is not provided")
	}
	if newUser.Password == "" {
		return nil, fmt.Errorf("user password is not provided")
	}

	return userService.userRepo.AddUser(newUser)
}

func (userSerice *UserServiceImpl) LoginUser(email string, password string) (*repo.User, error) {
	return userSerice.userRepo.AuthenticateUser(email, password)
}
