package services

import (
	"github.com/google/uuid"
	"wimf-services/dto"
	"wimf-services/entities"
	"wimf-services/helpers"
	"wimf-services/repositories"
)

type AuthService interface {
	Register(user dto.UserCreateDto) dto.UserDto
	Login(username, password string) interface{}
	IsDuplicateUsername(username string) bool
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepository repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepository,
	}
}

func (a *authService) Register(user dto.UserCreateDto) dto.UserDto {
	toCreate := entities.User{
		ID:       uuid.New().String(),
		Username: user.Username,
		Password: helpers.HashPassword(user.Password),
	}
	created := a.userRepository.Save(toCreate)

	return dto.UserDto{
		ID:       created.ID,
		Username: created.Username,
	}
}

func (a *authService) Login(username, password string) interface{} {
	res := a.userRepository.FindByUsername(username)

	if user, ok := res.(entities.User); ok {
		same := helpers.ComparePassword(user.Password, password)

		if user.Username == username && same {
			return res
		}
	}
	return false
}

func (a *authService) IsDuplicateUsername(username string) bool {
	res := a.userRepository.Exists(username)
	return res.Error == nil
}
