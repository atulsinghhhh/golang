package main

import (
	"fmt"
	"go-tweet/internal/config"
	"log"

	"github.com/gin-gonic/gin"
)

func main(){

	r:=gin.Default()

	cfg,err:=config.MustLoad()
	if err!=nil {
		log.Fatal(err)
	}
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	server:=fmt.Sprintf("localhost: %s",cfg.Port)
	r.Run(server)
}