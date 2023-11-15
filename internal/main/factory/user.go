package factory

import (
	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/usecase"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/cryptography"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/repository"
)

func CreateUserFactory() *usecase.CreateUserUsecase {
	return usecase.NewCreateUserUsecase(repository.NewUserRepository(), cryptography.NewBcryptAdapter())
}
