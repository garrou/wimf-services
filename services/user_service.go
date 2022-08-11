package services

import "wimf-services/repositories"

type UserService interface {
	Get(id string) interface{}
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

func (u *userService) Get(id string) interface{} {
	return u.userRepository.FindById(id)
}
