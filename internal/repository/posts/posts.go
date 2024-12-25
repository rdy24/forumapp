package posts

import (
	"context"
	"strings"

	"github.com/rdy24/forumapp/internal/model/posts"
	"github.com/rs/zerolog/log"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts(user_id, post_title, post_content, post_hashtags, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`
	_, err := r.db.ExecContext(ctx, query, model.UserId, model.PostTitle, model.PostContent, model.PostHashtags, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) GetAllPost(ctx context.Context, limit, offset int) (posts.GetAllPostResponse, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags FROM posts p JOIN users u ON p.user_id = u.id ORDER BY p.updated_at DESC LIMIT ? OFFSET ?`

	rows, err := r.db.QueryContext(ctx, query, limit, offset)

	response := posts.GetAllPostResponse{}

	if err != nil {
		return response, err
	}

	defer rows.Close()

	data := make([]posts.Post, 0)

	for rows.Next() {
		var (
			model    posts.PostModel
			username string
		)
		err := rows.Scan(&model.ID, &model.UserId, &username, &model.PostTitle, &model.PostContent, &model.PostHashtags)

		if err != nil {
			return response, err
		}

		data = append(data, posts.Post{
			ID:           model.ID,
			UserId:       model.UserId,
			Username:     username,
			PostTitle:    model.PostTitle,
			PostContent:  model.PostContent,
			PostHashtags: strings.Split(model.PostHashtags, ","),
		})
	}

	response.Data = data
	response.Pagination = posts.Pagination{
		Limit:  int64(limit),
		Offset: int64(offset),
	}

	return response, nil
}

func (r *repository) GetPostByID(ctx context.Context, id int64) (*posts.Post, error) {
	query := `SELECT p.id, p.user_id, u.username, p.post_title, p.post_content, p.post_hashtags, uv.is_liked FROM posts p JOIN users u ON p.user_id = u.id JOIN user_activities uv ON p.id = uv.post_id WHERE p.id = ?`

	row := r.db.QueryRowContext(ctx, query, id)

	var (
		model    posts.PostModel
		username string
		isLiked  bool
	)

	err := row.Scan(&model.ID, &model.UserId, &username, &model.PostTitle, &model.PostContent, &model.PostHashtags, &isLiked)

	if err != nil {
		log.Error().Err(err).Msg("error get post by id")
		return nil, err
	}

	return &posts.Post{
		ID:           model.ID,
		UserId:       model.UserId,
		Username:     username,
		PostTitle:    model.PostTitle,
		PostContent:  model.PostContent,
		PostHashtags: strings.Split(model.PostHashtags, ","),
		IsLiked:      isLiked,
	}, nil
}
