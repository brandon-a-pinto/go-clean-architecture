package usecase

import (
	"context"
	"time"

	errors "github.com/brandon-a-pinto/go-clean-architecture/internal/application/error"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/dto"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/protocol"
)

type AuthenticateUserUsecase struct {
	UserRepository protocol.IUserRepository
	BcryptAdapter  protocol.IBcryptAdapter
	JWTAdapter     protocol.IJWTAdapter
}

func NewAuthenticateUserUsecase(userRepository protocol.IUserRepository, bcryptAdapter protocol.IBcryptAdapter, jwtAdapter protocol.IJWTAdapter) *AuthenticateUserUsecase {
	return &AuthenticateUserUsecase{
		UserRepository: userRepository,
		BcryptAdapter:  bcryptAdapter,
		JWTAdapter:     jwtAdapter,
	}
}

func (u *AuthenticateUserUsecase) Exec(c context.Context, timeout time.Duration, input dto.AuthenticateUserInput) (*dto.AuthenticateUserOutput, error) {
	ctx, cancel := context.WithTimeout(c, timeout)
	defer cancel()

	user, err := u.UserRepository.FindByEmail(ctx, input.Email)
	if err != nil {
		return nil, err
	}

	err = u.BcryptAdapter.Compare(user.Password, input.Password)
	if err != nil {
		return nil, errors.NewUnauthorizedError(err)
	}

	token, err := u.JWTAdapter.Generate(user.ID.String())
	if err != nil {
		return nil, err
	}

	dto := &dto.AuthenticateUserOutput{
		AccessToken: token,
	}

	return dto, nil
}
