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
)