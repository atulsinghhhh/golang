package post

import (
	"context"
	"database/sql"
	"go-tweet/internal/model"
)

type PostRepo interface{
	StorePost(ctx context.Context, post *model.PostModel) (int64, error)
}

type postRepos struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) PostRepo {
	return &postRepos{
		db: db,
	}
}

