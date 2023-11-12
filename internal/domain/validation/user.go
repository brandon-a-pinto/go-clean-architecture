package validation

import (
	"errors"
	"regexp"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
)

func ValidateUser(input dto.CreateUserInput) error {
	if input.Email == "" {
		return errors.New("email is required")
	}
	if input.Username == "" {
		return errors.New("username is required")
	}
	if input.DisplayName == "" {
		return errors.New("display_name is required")
	}
	if input.Password == "" {
		return errors.New("password is required")
	}
	if len(input.Password) < 6 {
		return errors.New("password must be at least 6 characters")
	}
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(input.Email) {
		return errors.New("email is already taken")
	}

	return nil
}
