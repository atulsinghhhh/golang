package user

import (
	"context"
	"go-tweet/internal/model"
)

func (r *UserRepos) CreateUser(ctx context.Context, user *model.Usermodel) (int64, error) {
	// Implementation for creating user

	query:= `INSERT INTO users (email,username,password,created_at,updated_at) VALUES (?,?,?,?,?)`

	result,err:=r.db.ExecContext(ctx,query,user.Email,user.Username,user.Password,user.CreatedAt,user.UpdatedAt)
	if err!=nil{
		return 0,err
	}
	user_id,err:= result.LastInsertId()
	if err!=nil{
		return 0,err
	}
	return user_id,nil
}