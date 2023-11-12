package entity

import (
	"time"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/validation"
	"github.com/brandon-a-pinto/go-clean-architecture/pkg/domain/entity"
)

type User struct {
	ID          entity.ID
	Email       string
	Username    string
	DisplayName string
	Password    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewUser(input dto.CreateUserInput) (*User, error) {
	user := &User{
		ID:          entity.NewID(),
		Email:       input.Email,
		Username:    input.Username,
		DisplayName: input.DisplayName,
		Password:    input.Password,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	err := validation.ValidateUser(input)
	if err != nil {
		return nil, err
	}
	return user, nil
}
