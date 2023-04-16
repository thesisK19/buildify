package handler

import (
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("HealthCheck")
	// Send a success response
	Send(w, http.StatusOK, HealthCheckResponse{})
}

func Test(w http.ResponseWriter, r *http.Request) {
	log.Println("Test")
	// Send a success response
	Send(w, http.StatusOK, HealthCheckResponse{})
}
