package user

import (
	"go-tweet/internal/config"
	"go-tweet/internal/respository/user"
)

type UserService interface {
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
