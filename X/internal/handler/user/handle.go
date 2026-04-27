package user

import (
	"go-tweet/internal/services/user"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	api *gin.Engine
	userService  user.UserService
}

func NewHandler(api *gin.Engine,userService user.UserService) *Handler{
	return &Handler{
		api: api,
		userService: userService,
	}
}

func (h *Handler) RouteList(){
	
}