package posts

import (
	"context"

	"github.com/rdy24/forumapp/internal/model/posts"
)

func (r *repository) CreatePost(ctx context.Context, model posts.PostModel) error {
	query := `INSERT INTO posts (user_id, post_title, post_content, post_hastags, created_at, updated_at, ,created_by,updated_by) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserId, model.PostTitle, model.PostContent, model.PostHastags, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil

}
