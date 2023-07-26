package helper

import (
	"encoding/json"
	"net/http"
)

//Response comment
type Response struct {
	Message string `json:"message,omitempty"`
}

//RespondError template
func RespondError(w http.ResponseWriter, r *http.Request, status int, err error) {
	RespondMessage(w, r, status, err.Error())
}

//RespondMessage template
func RespondMessage(w http.ResponseWriter, r *http.Request, status int, message string) {
	RespondJSON(w, r, status, Response{Message: message})
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
