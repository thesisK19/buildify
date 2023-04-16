package handler

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	// Send a success response
	Send(w, http.StatusOK, "health check OK!!")
}
