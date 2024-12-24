package posts

import (
	"context"

	"github.com/rdy24/forumapp/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (s *service) GetPostByID(ctx context.Context, id int64) (*posts.GetPostResponse, error) {
	postDetail, err := s.postRepo.GetPostByID(ctx, id)

	if err != nil {
		log.Error().Err(err).Msg("error get post by id")
		return nil, err
	}

	likeCount, err := s.postRepo.CountLike(ctx, id)

	if err != nil {
		log.Error().Err(err).Msg("error count like")
		return nil, err
	}

	comments, err := s.postRepo.GetCommentByPostId(ctx, id)

	if err != nil {
		log.Error().Err(err).Msg("error get comment by post id")
		return nil, err
	}

	return &posts.GetPostResponse{
		PostDetail: posts.Post{
			ID:           postDetail.ID,
			UserId:       postDetail.UserId,
			Username:     postDetail.Username,
			PostTitle:    postDetail.PostTitle,
			PostContent:  postDetail.PostContent,
			PostHashtags: postDetail.PostHashtags,
			IsLiked:      postDetail.IsLiked,
		},
		LikeCount: likeCount,
		Comments:  comments,
	}, nil
}
