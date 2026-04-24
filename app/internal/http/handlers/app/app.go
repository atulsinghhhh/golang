package app

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"

	"github.com/atulsinghhhh/golang/internal/storage"
	"github.com/atulsinghhhh/golang/internal/types"
	"github.com/atulsinghhhh/golang/internal/utils/response"
)

func New(storage storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r*http.Request){

		slog.Info("creating my new app")
		var app types.App

		error:=json.NewDecoder(r.Body).Decode(&app)
		if errors.Is(error,io.EOF){
			response.WriteJson(w, http.StatusBadRequest,response.GeneralError)
			return
		}
		if error!=nil{
			response.WriteJson(w,http.StatusBadRequest,response.GeneralError(error))
			return
		}

		// request validation

		err:=validator.New().Struct(app)
		if err!=nil {
			validateErrors:=err.(validator.ValidationErrors)
			response.WriteJson(w,http.StatusBadRequest,response.ValidationError(validateErrors))
			return
		}




		response.WriteJson(w, http.StatusCreated,map[string]string{"success":"ok"})


		w.Write([]byte("Welcome to my app"))
	}
}