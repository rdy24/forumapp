package memberships

import (
	"context"
	"database/sql"

	"github.com/rdy24/forumapp/internal/model/memberships"
)

func (r *repository) GetUser(ctx context.Context, email, username string, userId int64) (*memberships.UserModel, error) {
	query := `SELECT id, email, username, password, created_at, updated_at, created_by, updated_by FROM users WHERE email = ? OR username = ? OR userId = ?`

	row := r.db.QueryRowContext(ctx, query, email, username, userId)

	var response memberships.UserModel
	err := row.Scan(&response.ID, &response.Email, &response.Username, &response.Password, &response.CreatedAt, &response.UpdatedAt, &response.CreatedBy, &response.UpdatedBy)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &response, nil
}

func (r *repository) CreateUser(ctx context.Context, model *memberships.UserModel) error {
	query := `INSERT INTO users (email, username, password, created_at, updated_at, created_by, updated_by) VALUES (?, ?, ?, ?, ?, ?, ?)`

	_, err := r.db.ExecContext(ctx, query, model.Email, model.Username, model.Password, model.CreatedAt, model.UpdatedAt, model.CreatedBy, model.UpdatedBy)

	if err != nil {
		return err
	}

	return nil
}
