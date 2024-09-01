package service

import (
	"dadandev.com/golang-dasar/internal/domain"
	"dadandev.com/golang-dasar/internal/dto"
)

type authService struct {
	userRepository domain.UserRepository
}

func NewAuth(repository domain.UserRepository) domain.UserService {
	return &authService{
		userRepository: repository,
	}
}
func (t authService) Authenticate() dto.ResponseUserToken {
	return dto.ResponseUserToken{
		Status:  200,
		Message: "Berhasil login",
	}
}
