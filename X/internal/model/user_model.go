package model

import "time"

type (
	Usermodel struct {
		ID        int64
		Email     string
		Username  string
		Password  string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	RefreshTokenModel struct {
		ID 	  int64
		UserID int64
		RefreshToken string
		ExpiresAt time.Time
		CreatedAt time.Time
		UpdatedAt time.Time

	}
)