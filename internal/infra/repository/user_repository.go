package repository

import (
	"context"
	"errors"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/entity"
	"github.com/brandon-a-pinto/go-clean-architecture/pkg/infra/database"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (id, email, username, display_name, password, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := database.DB.ExecContext(ctx, query, user.ID, user.Email, user.Username, user.DisplayName, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return errors.New("could not insert user into database")
	}
	return nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*entity.User, error) {
	var user entity.User
	row := database.DB.QueryRowContext(ctx, "SELECT id, email, username, display_name, password FROM users WHERE email=$1", email)
	if err := row.Scan(&user.ID, &user.Email, &user.Username, &user.DisplayName, &user.Password); err != nil {
		return nil, errors.New("email does not exist")
	}
	return &user, nil
}
