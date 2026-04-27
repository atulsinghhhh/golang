package user

import (
	"context"
	"go-tweet/internal/config"
	"go-tweet/internal/dto"
	"go-tweet/internal/respository/user"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterUserRequest) (int64, int, error)
	Login (ctx context.Context, req *dto.LoginUserRequest) (string,string, int, error)
}

type Service struct {
	cfg *config.Config
	userRepo user.UserRepo
}

func NewService(cfg *config.Config, userRepo user.UserRepo) UserService {
	return &Service{
		cfg: cfg,
		userRepo: userRepo,
	}
}
