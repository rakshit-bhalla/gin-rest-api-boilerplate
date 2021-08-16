package services

import (
	"rakshit.dev/gin-rest-api-boilerplate/errors"
	"rakshit.dev/gin-rest-api-boilerplate/models"
	"rakshit.dev/gin-rest-api-boilerplate/repositories"
)

type User = models.User
type UserReq = models.UserReq
type Error = errors.Error

// UserService ...
type UserService interface {
	GetUser(userID string) (*User, *Error)
	GetUsers() ([]User, *Error)
	DeleteUser(userID string) (*User, *Error)
	CreateUser(userReq UserReq) (*User, *Error)
}

type userService struct {
	userRepository repositories.UserRepository
}

// NewUserService ...
func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepository,
	}
}

func (u *userService) GetUser(userID string) (*User, *Error) {
	return u.userRepository.GetUser(userID)
}

func (u *userService) GetUsers() ([]User, *Error) {
	return u.userRepository.GetUsers()
}

func (u *userService) DeleteUser(userID string) (*User, *Error) {
	return u.userRepository.DeleteUser(userID)
}

func (u *userService) CreateUser(userReq UserReq) (*User, *Error) {
	return u.userRepository.CreateUser(userReq)
}
