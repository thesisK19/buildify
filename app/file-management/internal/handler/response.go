package handler

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Code    string         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Send(w http.ResponseWriter, res Response) error {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK) // alway return 200 Ok ?

	err := json.NewEncoder(w).Encode(res)
	return err
}
