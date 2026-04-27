package user

import (
	"context"
	"errors"
	"go-tweet/internal/dto"
	"go-tweet/internal/model"
	"go-tweet/sql/jwt"
	"go-tweet/sql/refreshtoken"
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) Login(ctx context.Context, req *dto.LoginUserRequest) (string,string,int,error){
	// check user exists
	userExists,err:=s.userRepo.GetUserByEmailorUsername(ctx,req.Email,"")
	if err!=nil {
		return "","",http.StatusInternalServerError,err
	}
	if userExists==nil {
		return "","",http.StatusBadRequest,errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExists.Password), []byte(req.Password))
	if err != nil {
		return "","",http.StatusUnauthorized,errors.New("invalid credentials")
	}

	token,err:=jwt.CreateToken(userExists.ID,userExists.Username,s.cfg.Secret_jwt)
	if err!=nil {
		return "","",http.StatusInternalServerError,err
	}

	now:= time.Now()
	refresh_token,err:=s.userRepo.GetRefreshToken(ctx,userExists.ID,now)
	if err!=nil {
		return "","",http.StatusInternalServerError,err
	}

	if refresh_token!=nil {
		return token,refresh_token.RefreshToken,http.StatusOK,nil
	}

	refresh,err:= refreshtoken.GenerateRefreshToken()
	if err!=nil {
		return "","",http.StatusInternalServerError,err
	}

	err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID: userExists.ID,
		RefreshToken: refresh,
		CreatedAt: now,
		UpdatedAt: now,
		ExpiresAt: now.Add(7*24*time.Hour),
	})
	if err!=nil {
		return "","",http.StatusInternalServerError,err
	}


	return token,refresh,http.StatusOK,nil
	// generate token
	// get refresh token
	// return token and refresh token

}