package domain

import "dadandev.com/golang-dasar/internal/dto"

type User struct {
	Nama string
}

type UserService interface {
	Authenticate() dto.ResponseUserToken
}

type UserRepository interface {
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
}
