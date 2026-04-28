package user

import (
	"context"
	"errors"
	"go-tweet/internal/dto"
	"go-tweet/internal/model"
	"go-tweet/sql/jwt"
	rt "go-tweet/sql/refreshtoken"
	"net/http"
	"time"
)

func (s *Service) RefreshToken(ctx context.Context, req *dto.RefreshTokenRequest, userID int64) (string, string, int, error) {
	// check user exist or not

	userexists,err:=s.userRepo.GetUserByID(ctx,userID)
	if err!=nil {
		return "", "", http.StatusInternalServerError, err
	}

	// get refresh token from database
	storedToken,err:=s.userRepo.GetRefreshToken(ctx,userID,time.Now())
	if err!=nil{
		return "","", http.StatusInternalServerError, errors.New("refresh token not found")
	}

	if req.RefreshToken != storedToken.RefreshToken {
		return "","", http.StatusUnauthorized, errors.New("invalid refresh token")
	}
	// check refresh token is match with re body

	token,err:=jwt.CreateToken(userID,userexists.Username,s.cfg.Secret_jwt)
	if err!=nil {
		return "","", http.StatusInternalServerError, err
	}

	err = s.userRepo.DeleteRefreshToken(ctx,userID)
	if err!=nil {
		return "","", http.StatusInternalServerError, err
	}
	refresh_token,err:= rt.GenerateRefreshToken()
	if err!=nil {
		return "","", http.StatusInternalServerError, err
	}

	now:= time.Now()
	err = s.userRepo.StoreRefreshToken(ctx, &model.RefreshTokenModel{
		UserID: userID,
		RefreshToken: refresh_token,
		CreatedAt: now,
		UpdatedAt: now,
		ExpiresAt: now.Add(7*24*time.Hour),
	})
	if err!=nil{
		return "","", http.StatusInternalServerError, err
	}
	// delete old refresh token & generate new access token & refresh token

	// store new refresh token into database
	return token, refresh_token, http.StatusOK, nil
}