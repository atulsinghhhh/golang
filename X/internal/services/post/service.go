package post

import (
	"context"
	"go-tweet/internal/config"
	"go-tweet/internal/dto"
	"go-tweet/internal/respository/post"
)

type PostService interface {
	CreatePost(ctx context.Context, req *dto.CreatePostRequest, userID int64) (int64, int, error)
}

type postService struct {
	cfg *config.Config
	postRepo post.PostRepo
}

func NewService(cfg *config.Config, postRepo post.PostRepo) PostService {
	return &postService{
		cfg: cfg,
		postRepo: postRepo,
	}
}