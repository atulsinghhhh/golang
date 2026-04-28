package post

import (
	"context"
	"go-tweet/internal/dto"
	"go-tweet/internal/model"
	"net/http"
	"time"
)

func (s *postService) CreatePost(ctx context.Context, req *dto.CreatePostRequest, userID int64) (int64, int, error) {

	now:=time.Now()

	insertedDB, err:= s.postRepo.StorePost(ctx,&model.PostModel{
		UserID: userID,
		Content: req.Content,
		CreatedAt: now,
		UpdatedAt: now,
	})
	if err!=nil {
		return 0,http.StatusInternalServerError,err
	}

	return insertedDB,http.StatusCreated,nil

}