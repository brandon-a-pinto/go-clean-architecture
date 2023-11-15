package factory

import (
	"database/sql"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/application/usecase"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/cryptography"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/repository"
)

func CreateUserFactory(db *sql.DB) *usecase.CreateUserUsecase {
	return usecase.NewCreateUserUsecase(repository.NewUserRepository(db), cryptography.NewBcryptAdapter())
}
