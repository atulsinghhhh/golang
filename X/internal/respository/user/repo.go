package user

import (
	"context"
	"database/sql"
	"go-tweet/internal/model"
)

type UserRepo interface {
	GetUserByEmailorUsername(ctx context.Context,email,username string) (*model.Usermodel, error)
	CreateUser(ctx context.Context, user *model.Usermodel) (int64, error)
}

type UserRepos struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) UserRepo {
	return &UserRepos{
		db: db,
	}
}