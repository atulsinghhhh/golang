package post

import (
	"go-tweet/internal/dto"
	"net/http"
	"github.com/gin-gonic/gin"
)

func (h *Handler) CreatePost (c *gin.Context) {

	var (
		ctx = c.Request.Context()
		req dto.CreatePostRequest
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
	userId:= c.GetInt64("user_id")
	post_id, status, err := h.postService.CreatePost(ctx, &req, userId)
	if err != nil {
		c.JSON(status, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"post_id": post_id,
	})
}