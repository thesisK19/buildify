package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/thesisK19/buildify/app/file-mgt/internal/constant"
)

type UploadImageResponse struct {
	Code    constant.Code `json:"code"`
	Message string        `json:"message"`
	Url     string        `json:"url"`
}

func Send(w http.ResponseWriter, statusCode int, res interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(statusCode)

	if err := json.NewEncoder(w).Encode(res); err != nil {
		// Log the error
		log.Printf("Error encoding response: %s", err.Error())
	}
}
