package utils

import (
	"time"

	"github.com/golang-jwt/jwt"
)

type TokenParams struct {
	Ttl time.Time
	IsAdmin bool
	Name string
	Payload interface{}
	PrivateKey string
}

func CreateToken(tokenParam TokenParams) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = tokenParam.Name
	claims["admin"] = tokenParam.IsAdmin
	claims["sub"] = tokenParam.Payload
	claims["exp"] = tokenParam.Ttl.Unix()

	tokenString, err := token.SignedString([]byte(tokenParam.PrivateKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil;
}