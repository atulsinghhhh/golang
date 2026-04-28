package main

import (
	"fmt"
	"go-tweet/internal/config"
	userhandler "go-tweet/internal/handler/user"
	userrepo "go-tweet/internal/respository/user"
	userservice "go-tweet/internal/services/user"

	postrepo "go-tweet/internal/respository/post"
	postservice "go-tweet/internal/services/post"
	posthandler "go-tweet/internal/handler/post"
	"go-tweet/sql/internalsql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {

	r := gin.Default()
	validate:=validator.New()

	cfg, err := config.MustLoad()
	if err != nil {
		log.Fatal(err)
	}
	db, error := internalsql.ConnectMySQL(cfg)
	if error != nil {
		log.Fatal(error)
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})
	userRepo := userrepo.NewRepo(db)
	postRepo := postrepo.NewRepo(db)
	userService := userservice.NewService(cfg, userRepo)
	postService := postservice.NewService(cfg, postRepo)
	userHandler := userhandler.NewHandler(r, validate, userService)
	postHandler := posthandler.NewHandler(r, validate, postService)
	userHandler.RouteList(cfg.Secret_jwt)
	postHandler.RouteList(cfg.Secret_jwt)

	server := fmt.Sprintf("localhost: %s", cfg.Port)
	r.Run(server)
}
