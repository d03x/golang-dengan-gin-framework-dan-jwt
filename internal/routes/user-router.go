package routes

import (
	"time"

	"dadandev.com/golang-dasar/internal/domain"
	"dadandev.com/golang-dasar/internal/dto"
	"dadandev.com/golang-dasar/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type userRouter struct {
	userService domain.UserService
}

func NewUser(gin *gin.Engine, userService domain.UserService) {
	user := userRouter{
		userService: userService,
	}
	gin.POST("/api/v1/login", user.Login)
}

func (ur userRouter) Login(context *gin.Context) {
	_ = ur.userService.Authenticate()
	var dtoUserResponse dto.ResponseAcessToken
	_, err := utils.GenerateJWTToken(jwt.MapClaims{
		"user_id": "2804",
		"expired": time.Now().Add(time.Hour * 1).Unix(),
	}, &dtoUserResponse)
	if err != nil {
		context.Error(err)
	}
	// AcessToken, _ := utils.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoxNzI1NDMyOTMzLCJ1c2VyX2lkIjoiMjgwNCJ9.7ADF_eeiN-Mu1uLe8Sen_XsRV9e_CndfBHUyqMkc6lA")
	var userDto dto.UserRequestDto

	context.ShouldBindBodyWithJSON(&userDto)
	context.JSON(200, dtoUserResponse)
}
