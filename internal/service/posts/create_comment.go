package posts

import (
	"context"
	"strconv"
	"time"

	"github.com/rdy24/forumapp/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) CreateComment(ctx context.Context, postId, userId int64, req posts.CreateCommentRequest) error {
	model := posts.CommentModel{
		PostId:         postId,
		UserId:         userId,
		CommentContent: req.CommentContent,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		CreatedBy:      strconv.FormatInt(userId, 10),
		UpdatedBy:      strconv.FormatInt(userId, 10),
	}

	err := s.postRepo.CreateComment(ctx, model)
	if err != nil {
		log.Error().Err(err).Msg("failed to create comment")
	}

	return nil
}
