package user

import (
	"go-tweet/internal/dto"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) RefreshToken(c *gin.Context) {

	var (
		ctx = c.Request.Context()
		req dto.RefreshTokenRequest
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}
	if err:= h.validate.Struct(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID:=c.GetInt64("user_id")
	token, refresh_token, status, err := h.userService.RefreshToken(ctx, &req, userID)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"token": token,
		"refresh_token": refresh_token,
	})
}