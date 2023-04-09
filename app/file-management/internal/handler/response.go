package handler

import (
	"encoding/json"
	"net/http"

	"github.com/thesisK19/buildify/app/file-management/internal/constant"
)

type Response struct {
	Code    constant.Code `json:"code"`
	Message string        `json:"message"`
	Data    interface{}   `json:"data"`
}

func Send(w http.ResponseWriter, res Response) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK) // alway return 200 Ok ?

	err := json.NewEncoder(w).Encode(res)
	return err
}
