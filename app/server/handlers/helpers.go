package handlers

import (
	"fmt"
	"net/http"

	"app/logs"
)

const (
	contentType = "Content-Type"
	authToken   = "X-Auth-Token"
	jsonContent = "application/json"
	mp3Content  = "audio/mpeg"
)

func logRequest(r *http.Request) {
	logs.Info("Received request: %s", requestSummary(r))
}

func requestSummary(r *http.Request) string {
	return fmt.Sprintf("%v %v", r.Method, r.URL)
}

func BadRequest(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "400 bad request", http.StatusBadRequest)
}

func Unauthorized(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "401 unauthorized", http.StatusUnauthorized)
}

func Forbidden(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "403 forbidden", http.StatusForbidden)
}

func Conflict(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "409 conflict", http.StatusConflict)
}

func InternalServerError(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "500 internal server error", http.StatusInternalServerError)
}

func NotFound(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "404 Not Found", http.StatusNotFound)
}
