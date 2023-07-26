package server

import (
	"context"
	"cron-test/pkg/config"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/robfig/cron"
)

//Server info
type Server struct {
	Port int
	Mux  *mux.Router
	Cjob *cron.Cron
}

// NewServer ...
func NewServer(conf config.Config) *Server {
	location, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println("Error loading time location. using default.")
		location = time.FixedZone("UTC", -6*3600)
	}

	server := Server{
		Mux:  mux.NewRouter(),
		Port: conf.Port,
		Cjob: cron.NewWithLocation(location),
	}

	return &server
}

//Run ...
func (s *Server) Run() error {
	fmt.Printf("Starting Server on :%d \n", s.Port)
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: s.Mux,
	}
	// s.Cjob.Start()

	//Graceful shutdown
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		fmt.Println("Shutting Down Server")

		if err := srv.Shutdown(context.Background()); err != nil {
			fmt.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server ListenAndServe: %v", err)
	}
	return nil
}
