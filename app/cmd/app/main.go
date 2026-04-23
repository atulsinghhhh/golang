package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/atulsinghhhh/golang/internal/config"
	"github.com/atulsinghhhh/golang/internal/http/handlers/app"
)

func main() {
	fmt.Println("welcome to golang project")

	/*
		load config
		db setup
		setup router
		setp server
	
	*/

	cfg:=config.MustLoad()
	router:=http.NewServeMux()
	router.HandleFunc("POST /api/v1/app",app.New())

	server:=http.Server {
		Addr: cfg.HTTPServer.Addr,
		Handler: router,
	}
	fmt.Println("server started")

	done:=make(chan os.Signal,1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		error:=server.ListenAndServe()
		if error!=nil{
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("shutting down server")


	
	ctx,cancel:=context.WithTimeout(context.Background(),5*time.Second)
	defer cancel()

	error:=server.Shutdown(ctx)
	if error!=nil{
		slog.Error("failed to setup server",slog.String("error",error.Error()))
	}

	slog.Info("server shutdown successfully")

	// fmt.Println("server started")

}