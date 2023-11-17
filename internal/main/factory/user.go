package factory

import (
	"github.com/brandon-a-pinto/go-clean-architecture/configs"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/usecase"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/cryptography"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/repository"
)

func CreateUserFactory() *usecase.CreateUserUsecase {
	return usecase.NewCreateUserUsecase(repository.NewUserRepository(), cryptography.NewBcryptAdapter())
}

func AuthenticateUserFactory() *usecase.AuthenticateUserUsecase {
	config := configs.LoadConfig()
	return usecase.NewAuthenticateUserUsecase(
		repository.NewUserRepository(),
		cryptography.NewBcryptAdapter(),
		cryptography.NewJWTAdapter([]byte(config.JWTSecret), config.JWTExpiresIn),
	)
}
