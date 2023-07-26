package server

import (
	"fmt"
	"net/http"

	"github.com/robfig/cron"
)

//Routes routes for server
func (s *Server) Routes() {
	s.Mux.HandleFunc("/health", health()).Methods("GET")
	s.Mux.HandleFunc("/stopCron", stopCron(s.Cjob)).Methods("GET")
	s.Mux.HandleFunc("/startCron", startCron(s.Cjob)).Methods("GET")
	s.Mux.HandleFunc("/infoCron", infoCron(s.Cjob)).Methods("GET")
	s.Mux.HandleFunc("/runCron", runCron(s.Cjob)).Methods("GET")
}

func health() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		RespondStatus(w, r, http.StatusOK)
	}
}

func stopCron(c *cron.Cron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Stopping Scheduler")
		c.Stop()
		RespondStatus(w, r, http.StatusOK)
	}
}
func startCron(c *cron.Cron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Start Scheduler")
		c.Start()
		RespondStatus(w, r, http.StatusOK)
	}
}
func infoCron(c *cron.Cron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Scheduler info:")
		e := c.Entries()
		body := ""
		for _, v := range e {
			body += fmt.Sprintf("next Run: %s\n", v.Next)
		}
		RespondText(w, r, http.StatusOK, body)
	}
}
func runCron(c *cron.Cron) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c.Run()
		RespondStatus(w, r, http.StatusOK)
	}
}
