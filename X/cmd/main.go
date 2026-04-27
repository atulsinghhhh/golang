package main

import (
	"fmt"
	"go-tweet/internal/config"
	userhandler "go-tweet/internal/handler/user"
	userrepo "go-tweet/internal/respository/user"
	userservice "go-tweet/internal/services/user"
	"go-tweet/sql/internalsql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

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
	userService := userservice.NewService(cfg, userRepo)
	userHandler := userhandler.NewHandler(r, userService)
	userHandler.RouteList()

	server := fmt.Sprintf("localhost: %s", cfg.Port)
	r.Run(server)
}
