package dto

type (
	RegisterUserRequest struct {
		Email    string `json:"email"  validate:"email"`
		Password string `json:"password" validate:"min=6"`
		Username string `json:"username" validate:"required"`
	}

	RegisterUserResponse struct {
		ID int64 `json:"id"`
	}
)

type (
	LoginUserRequest struct {
		Email string `json:"email"  validate:"email"`
		Password string `json:"password" validate:"min=6"`
	}
	LoginUserResponse struct {
		Token string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}
)
