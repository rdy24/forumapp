package posts

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/rdy24/forumapp/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) UpsertUserActivity(ctx context.Context, postId, userId int64, req posts.UserActivityRequest) error {

	model := posts.UserActivityModel{
		PostId:    postId,
		UserId:    userId,
		IsLiked:   req.IsLiked,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: strconv.FormatInt(userId, 10),
		UpdatedBy: strconv.FormatInt(userId, 10),
	}

	userActivity, err := s.postRepo.GetUserActivity(ctx, model)

	if err != nil {
		log.Error().Err(err).Msg("failed to get user activity")
	}

	if userActivity == nil {
		if !req.IsLiked {
			return errors.New("anda belum pernah menyukai post ini")
		}

		err = s.postRepo.CreateUserActivity(ctx, model)

	} else {
		err = s.postRepo.UpdateUserActivity(ctx, model)
	}

	if err != nil {
		log.Error().Err(err).Msg("failed to upsert user activity")
		return err
	}

	return nil
}
