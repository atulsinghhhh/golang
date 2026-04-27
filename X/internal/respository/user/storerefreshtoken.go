package user

import (
	"context"
	"go-tweet/internal/model"
)

func (s *UserRepos) StoreRefreshToken(ctx context.Context, token *model.RefreshTokenModel) error {
	query:= `INSERT INTO refresh_tokens (user_id, refresh_token, expires_at) VALUES (?, ?, ?)`

	_, err := s.db.ExecContext(ctx, query, token.UserID, token.RefreshToken, token.ExpiresAt)
	if err != nil {
		return err
	}
	return err
}