package posts

import (
	"context"

	"github.com/rdy24/forumapp/internal/configs"
	"github.com/rdy24/forumapp/internal/model/posts"
)

type postRepository interface {
	CreatePost(ctx context.Context, model posts.PostModel) error
	CreateComment(ctx context.Context, model posts.CommentModel) error

	CreateUserActivity(ctx context.Context, model posts.UserActivityModel) error
	GetUserActivity(ctx context.Context, model posts.UserActivityModel) (*posts.UserActivityModel, error)
	UpdateUserActivity(ctx context.Context, model posts.UserActivityModel) error
}

type service struct {
	postRepo postRepository
	cfg      *configs.Config
}

func NewService(cfg *configs.Config, postRepo postRepository) *service {
	return &service{
		cfg:      cfg,
		postRepo: postRepo,
	}
}
