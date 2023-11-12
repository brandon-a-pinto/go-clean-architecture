package protocol

import (
	"context"

	"github.com/brandon-a-pinto/go-clean-architecture/internal/domain/entity"
)

type IUserRepository interface {
	Save(ctx context.Context, user *entity.User) error
}
