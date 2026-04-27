package user

import (
	"go-tweet/internal/dto"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) Resgister(c *gin.Context) {

	var (
		ctx = c.Request.Context()
		req dto.RegisterUserRequest
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
	user_id, status, err := h.userService.Register(ctx, &req)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"user_id": user_id,
	})
}