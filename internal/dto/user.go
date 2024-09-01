package dto

type ResponseUserToken struct {
	Status  int
	Message string
}
type UserRequestDto struct {
	Email      string `form:"email"`
	Password   string `form:"password"`
	RememberMe bool   `form:"remember_me"`
}
