package user

import (
	"context"
	"database/sql"
	"go-tweet/internal/model"
	"time"
)

type UserRepo interface {
	GetUserByEmailorUsername(ctx context.Context,email,username string) (*model.Usermodel, error)
	CreateUser(ctx context.Context, user *model.Usermodel) (int64, error)
	GetRefreshToken(ctx context.Context,userID int64, now time.Time) (*model.RefreshTokenModel, error)
	StoreRefreshToken(ctx context.Context, token *model.RefreshTokenModel) error
	GetUserByID(ctx context.Context, userID int64) (*model.Usermodel, error)
	DeleteRefreshToken(ctx context.Context, userID int64) error
}

type UserRepos struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) UserRepo {
	return &UserRepos{
		db: db,
	}
}