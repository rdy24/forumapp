package memberships

import (
	"context"
	"database/sql"
	"time"

	"github.com/rdy24/forumapp/internal/model/memberships"
)

func (r *repository) InsertRefershToken(ctx context.Context, model memberships.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens (user_id, refresh_token, expired_at, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.UserId, model.RefreshToken, model.ExpiredAt, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}

func (r *repository) GetRefreshToken(ctx context.Context, userId int64, now time.Time) (*memberships.RefreshTokenModel, error) {
	query := `SELECT * FROM refresh_tokens WHERE user_id = ? AND expired_at >= ?`

	var response memberships.RefreshTokenModel

	row := r.db.QueryRowContext(ctx, query, userId, now)

	err := row.Scan(&response.ID, &response.UserId, &response.RefreshToken, &response.ExpiredAt, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}
