package cryptography

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTAdapter struct {
	Secret    []byte
	ExpiresIn int
}

func NewJWTAdapter(secret []byte, exp int) *JWTAdapter {
	return &JWTAdapter{
		Secret:    secret,
		ExpiresIn: exp,
	}
}

func (j *JWTAdapter) Generate(id string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = id
	claims["exp"] = time.Now().Add(time.Minute * time.Duration(j.ExpiresIn)).Unix()

	tokenString, err := token.SignedString(j.Secret)
	if err != nil {
		return "", errors.New("could not generate jwt token")
	}

	return tokenString, nil
}
