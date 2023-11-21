package usecase

import (
	"context"
	"time"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/helper"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/entity"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/protocol"
)

type CreateUserUsecase struct {
	userRepository protocol.IUserRepository
	bcryptAdapter  protocol.IBcryptAdapter
}

func NewCreateUserUsecase(userRepository protocol.IUserRepository, bcryptAdapter protocol.IBcryptAdapter) *CreateUserUsecase {
	return &CreateUserUsecase{
		userRepository: userRepository,
		bcryptAdapter:  bcryptAdapter,
	}
}

func (u *CreateUserUsecase) Exec(c context.Context, timeout time.Duration, input dto.CreateUserInput) (*dto.CreateUserOutput, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	user, err := entity.NewUser(input)
	if err != nil {
		return nil, helper.NewBadRequestError(err)
	}

	hashedPassword, err := u.bcryptAdapter.Hash(user.Password, 12)
	if err != nil {
		return nil, err
	}
	user.Password = hashedPassword

	err = u.userRepository.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	dto := &dto.CreateUserOutput{
		Email:       user.Email,
		Username:    user.Username,
		DisplayName: user.DisplayName,
	}

	return dto, nil
}
