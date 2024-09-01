package repositories

import "dadandev.com/golang-dasar/internal/domain"

type userRepository struct {
}

func NewUser() domain.UserRepository {
	return &userRepository{}
}

func (userrepo userRepository) FindByEmail(email string) (user domain.User, err error) {

	user = domain.User{
		Nama: "Dadan hidayat",
	}
	return
}

func (userrepo userRepository) FindById(id int) (user domain.User, err error) {
	user = domain.User{
		Nama: "Dadan hidayat",
	}
	return
}
