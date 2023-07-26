package server

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Err string `json:"error,omitempty"`
}

func RespondError(w http.ResponseWriter, r *http.Request, status int, err error) {
	RespondJSON(w, r, status, ErrorResponse{Err: err.Error()})
}

//RespondJSON template
func RespondJSON(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// RespondStatus ...
func RespondStatus(w http.ResponseWriter, r *http.Request, status int) {
	Respond(w, r, status, "text/html", []byte(http.StatusText(status)))

}

// RespondText ...
func RespondText(w http.ResponseWriter, r *http.Request, status int, body string) {
	Respond(w, r, status, "text/html", []byte(body))

}

// Respond ...
func Respond(w http.ResponseWriter, r *http.Request, status int, contentType string, body []byte) {
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(status)
	w.Write(body)
}
