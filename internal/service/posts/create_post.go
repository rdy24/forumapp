package posts

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/rdy24/forumapp/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreatePost(ctx context.Context, userId int64, req posts.CreatePostRequest) error {
	postHastags := strings.Join(req.PostHastags, ",")

	model := posts.PostModel{
		UserId:      userId,
		PostTitle:   req.PostTitle,
		PostContent: req.PostContent,
		PostHastags: postHastags,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		CreatedBy:   strconv.FormatInt(userId, 10),
		UpdatedBy:   strconv.FormatInt(userId, 10),
	}

	err := s.postRepo.CreatePost(ctx, model)

	if err != nil {
		log.Error().Err(err).Msg("failed to create post")
		return err
	}

	return nil
}
