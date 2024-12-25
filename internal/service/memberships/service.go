package memberships

import (
	"context"
	"time"

	"github.com/rdy24/forumapp/internal/configs"
	"github.com/rdy24/forumapp/internal/model/memberships"
)

type membershipRepository interface {
	GetUser(ctx context.Context, email, username string, userID int64) (*memberships.UserModel, error)
	CreateUser(ctx context.Context, model *memberships.UserModel) error
	GetRefreshToken(ctx context.Context, userId int64, now time.Time) (*memberships.RefreshTokenModel, error)
	InsertRefershToken(ctx context.Context, model memberships.RefreshTokenModel) error
}

type service struct {
	membershipRepo membershipRepository
	cfg            *configs.Config
}

func NewService(cfg *configs.Config, membershipRepo membershipRepository) *service {
	return &service{
		cfg:            cfg,
		membershipRepo: membershipRepo,
	}
}
