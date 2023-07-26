package server

import (
	"fmt"
	"net/http"
	"time"
)

//Routes routes for server
func (s *Server) Routes() {
	s.Mux.HandleFunc("/health", health()).Methods("GET")
	s.Mux.HandleFunc("/ready", ready()).Methods("GET")
	s.Mux.HandleFunc("/sleep", s.sleep()).Methods("GET")
	s.Mux.HandleFunc("/isSleep", s.isSleep()).Methods("GET")
}

func health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		RespondStatus(w, r, http.StatusOK)
	}
}

func ready() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		one := 1
		if one == 1 {
			RespondStatus(w, r, http.StatusBadRequest)
			return
		}
		RespondStatus(w, r, http.StatusOK)
	}
}

func (s *Server) sleep() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.IsSleep = true
		fmt.Println("Sleeping")
		time.Sleep(30 * time.Second)
		s.IsSleep = false
		fmt.Println("Not Sleeping")
		RespondStatus(w, r, http.StatusOK)
	}
}

func (s *Server) isSleep() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if s.IsSleep {
			RespondText(w, r, http.StatusOK, "Yes, Im Sleeping")
			return
		}
		RespondText(w, r, http.StatusOK, "Im Awake")
	}
}
