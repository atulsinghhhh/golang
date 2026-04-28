package user

import (
	"context"
	"go-tweet/internal/model"
)

func (r *UserRepos) GetUserByID(ctx context.Context, userID int64) (*model.Usermodel, error) {
	query := `SELECT id, email,username,created_at,updated_at FROM users WHERE id = ?`

	row := r.db.QueryRowContext(ctx, query, userID)

	var result model.Usermodel
	err := row.Scan(&result.ID, &result.Email, &result.Username, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
