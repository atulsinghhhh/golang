package user

import (
	"go-tweet/internal/dto"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(c *gin.Context) {

	var (
		ctx = c.Request.Context()
		req dto.LoginUserRequest
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
	token, refresh_token, status, err := h.userService.Login(ctx, &req)
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