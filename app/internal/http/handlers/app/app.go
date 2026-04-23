package app

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/atulsinghhhh/golang/internal/types"
	"github.com/atulsinghhhh/golang/internal/utils/response"
	
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r*http.Request){

		var app types.App

		error:=json.NewDecoder(r.Body).Decode(&app)
		if errors.Is(error,io.EOF){
			response.WriteJson(w, http.StatusBadRequest,response.GeneralError)
			return
		}
		slog.Info("creating my new app")
		response.WriteJson(w, http.StatusCreated,map[string]string{"success":"ok"})


		w.Write([]byte("Welcome to my app"))
	}
}