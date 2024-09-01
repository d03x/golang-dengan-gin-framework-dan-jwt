package routes

import (
	"dadandev.com/golang-dasar/internal/domain"
	"dadandev.com/golang-dasar/internal/dto"
	"dadandev.com/golang-dasar/internal/utils"
	"github.com/gin-gonic/gin"
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

	// signedToken, err := utils.GenerateJWTToken(jwt.MapClaims{
	// 	"user_id": "2804",
	// 	"expired": time.Now().Add(time.Hour * 1).Unix(),
	// })
	// if err != nil {
	// 	context.Error(err)
	// }
	signedToken, _ := utils.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVkIjoxNzI1MjA2MzI2LCJ1c2VyX2lkIjoiMjgwNCJ9.Q9T6hCnhlbAJSFjFJ3VAYVRqQTxwLg5n3-tRjwDHOVs")
	var userDto dto.UserRequestDto
	context.ShouldBindBodyWithJSON(&userDto)
	context.JSON(200, signedToken)
}
