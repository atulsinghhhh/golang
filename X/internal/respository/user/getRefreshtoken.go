package user

import (
	"context"
	"go-tweet/internal/model"
	"time"
)

func (r *UserRepos) GetRefreshToken(ctx context.Context, userID int64, now time.Time) (*model.RefreshTokenModel, error) {
	query:= `SELECT id, user_id, refresh_token, expires_at
	FROM refresh_tokens
	WHERE user_id=? AND expires_at > ?`

	row:= r.db.QueryRowContext(ctx,query,userID,now)

	var result model.RefreshTokenModel
	err:=row.Scan(&result.ID,&result.UserID,&result.RefreshToken,&result.ExpiresAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}