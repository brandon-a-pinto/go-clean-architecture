package repository

import (
	"context"
	"database/sql"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/entity"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, user *entity.User) error {
	query := `INSERT INTO users (id, email, username, display_name, password, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := r.DB.ExecContext(ctx, query, user.ID, user.Email, user.Username, user.DisplayName, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}
	return nil
}
