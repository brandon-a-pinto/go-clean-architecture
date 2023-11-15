package repository

import (
	"context"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/entity"
	"github.com/brandon-a-pinto/go-clean-architecture/internal/infra/database"
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
		return err
	}
	return nil
}
