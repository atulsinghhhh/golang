package user

import (
	"context"
	"go-tweet/internal/model"
)

func (r *UserRepos) GetUserByEmailorUsername(ctx context.Context, email, username string) (*model.Usermodel, error) {
	query:= `SELECT id,username,email,password,created_at,updated_at
		FROM users
		WHERE email=? OR username=?
	`

	row:= r.db.QueryRowContext(ctx,query,email,username)

	var result model.Usermodel
	err:=row.Scan(&result.ID,&result.Username,&result.Email,&result.Password,&result.CreatedAt,&result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &result, nil
}