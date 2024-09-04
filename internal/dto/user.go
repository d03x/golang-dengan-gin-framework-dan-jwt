package dto

type ResponseUserToken struct {
	Status  int
	Message string
}
type UserRequestDto struct {
	Email      string `json:"email"`
	Password   string `json:"password"`
	RememberMe bool   `json:"remember_me"`
}

type ResponseAcessToken struct {
	AccessToken string `json:"access_token"`
}
