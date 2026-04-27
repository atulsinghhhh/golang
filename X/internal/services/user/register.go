package user

import (
	"context"
	"database/sql"
	"errors"
	"go-tweet/internal/dto"
	"go-tweet/internal/model"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)


func (s *Service) Register(ctx context.Context, req *dto.RegisterUserRequest) (int64, int, error){
	// check user already exists

	userexist,err:=s.userRepo.GetUserByEmailorUsername(ctx,req.Email,req.Username)
	if err!=nil && !errors.Is(err, sql.ErrNoRows){
		return 0,http.StatusInternalServerError,err
	}

	if userexist!=nil {
		return 0,http.StatusBadRequest,errors.New("user already exists")
	}

	passwordHash,err:=bcrypt.GenerateFromPassword([]byte(req.Password),bcrypt.DefaultCost)
	if err!=nil{
		return 0,http.StatusInternalServerError,err
	}

	now:=time.Now()
	userModel:=&model.Usermodel{
		Email: req.Email,
		Username: req.Username,
		Password: string(passwordHash),
		CreatedAt: now,
		UpdatedAt: now,
	}
	user_id,err:= s.userRepo.CreateUser(ctx, userModel)
	if err!=nil{
		return 0,http.StatusInternalServerError,err
	}
	return user_id,http.StatusCreated,nil
}

