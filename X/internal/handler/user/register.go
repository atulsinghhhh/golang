package user

import (
	"go-tweet/internal/dto"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Resgister(c *gin.Context) {

	var (
		ctx = c.Request.Context()
		req dto.RegisterUserRequest
	)
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": "invalid request",
		})
		return
	}
	user_id, status, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(201, gin.H{
		"user_id": user_id,
	})
}