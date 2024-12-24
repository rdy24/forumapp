package posts

import (
	"context"

	"github.com/rdy24/forumapp/internal/model/posts"
)

func (r *repository) CreateComment(ctx context.Context, model posts.CommentModel) error {
	query := `INSERT INTO comments (post_id, user_id, comment_content, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.PostId, model.UserId, model.CommentContent, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetCommentByPostId(ctx context.Context, postId int64) ([]posts.Comment, error) {
	query := `SELECT c.id, c.user_id, c.comment_content, u.username FROM comments c JOIN users u ON c.user_id = u.id WHERE post_id = ?`

	rows, err := r.db.QueryContext(ctx, query, postId)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	data := make([]posts.Comment, 0)

	for rows.Next() {
		var (
			comment  posts.Comment
			username string
		)

		err := rows.Scan(&comment.Id, &comment.UserId, &comment.CommentContent, &username)

		if err != nil {
			return nil, err
		}

		data = append(data, posts.Comment{
			Id:             comment.Id,
			UserId:         comment.UserId,
			CommentContent: comment.CommentContent,
			Username:       username,
		})
	}

	return data, nil
}
