package utils

import (
	"fmt"
	"time"

	"dadandev.com/golang-dasar/internal/dto"
	"github.com/dgrijalva/jwt-go"
)

var SignedKey = []byte("n983274hr98ew7r379847324")

func GenerateJWTToken(claims jwt.Claims, userResponse *dto.ResponseAcessToken) (signedToken string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(SignedKey)
	userResponse.AccessToken = signedToken
	return
}
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return SignedKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["expired"].(float64); ok {
			expTime := time.Unix(int64(exp), 0)
			if expTime.Before(time.Now()) {
				return claims, fmt.Errorf("token expired")
			}
		}
		return claims, nil
	} else {
		return nil, fmt.Errorf("invalid token")
	}
}
