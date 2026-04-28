package post

import (
	"context"
	"go-tweet/internal/model"
)

func (r *postRepos) StorePost(ctx context.Context, post *model.PostModel) (int64, error) {
	query:= `INSERT INTO posts (user_id,content,created_at,updated_at) VALUES (?,?,?,?)`

	result,err:=r.db.ExecContext(ctx,query,post.UserID,post.Content,post.CreatedAt,post.UpdatedAt)
	if err!=nil {
		return 0,err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}