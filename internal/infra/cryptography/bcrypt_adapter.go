package cryptography

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type BcryptAdapter struct{}

func NewBcryptAdapter() *BcryptAdapter {
	return &BcryptAdapter{}
}

func (b *BcryptAdapter) Hash(password string, salt int) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), salt)
	if err != nil {
		return "", errors.New("could not hash password")
	}
	return string(hash), nil
}

func (b *BcryptAdapter) Compare(hashedPassword, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return errors.New("incorrect password")
	}
	return nil
}
