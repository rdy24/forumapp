package posts

import (
	"context"

	"github.com/rdy24/forumapp/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetAllPost(ctx context.Context, pageSize, pageIndex int) (posts.GetAllPostResponse, error) {
	limit := pageSize
	offset := (pageIndex - 1) * pageSize

	response, err := s.postRepo.GetAllPost(ctx, limit, offset)

	if err != nil {
		log.Error().Err(err).Msg("error getting all post")
		return response, err
	}

	return response, nil
}
