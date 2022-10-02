package services

import (
	"wimf-services/dto"
	"wimf-services/entities"
	"wimf-services/helpers"
	"wimf-services/repositories"
)

type UserService interface {
	Get(id string) interface{}
	UpdateUsername(dto dto.UsernameDto) interface{}
	UpdatePassword(dto dto.PasswordDto) interface{}
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

func (u *userService) UpdateUsername(dto dto.UsernameDto) interface{} {
	res := u.userRepository.FindById(dto.UserId)

	if user, ok := res.(entities.User); ok {
		user.Username = dto.Username
		return u.userRepository.Save(user)
	}
	return nil
}

func (u *userService) UpdatePassword(dto dto.PasswordDto) interface{} {
	res := u.userRepository.FindById(dto.UserId)

	if user, ok := res.(entities.User); ok {
		same := helpers.ComparePassword(user.Password, dto.Current)

		if same {
			user.Password = helpers.HashPassword(dto.Password)
			return u.userRepository.Save(user)
		}
	}
	return nil
}
