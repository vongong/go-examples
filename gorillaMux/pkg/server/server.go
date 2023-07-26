package server

import (
	"context"
	"fmt"
	"gmux/pkg/config"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

type Server struct {
	Port    int
	Mux     *mux.Router
	IsSleep bool
}

func NewServer(conf config.Config) *Server {
	server := Server{
		Mux: mux.NewRouter(),
	}
	server.Port = conf.Port

	return &server
}

func (s *Server) Run() error {
	fmt.Println("Starting Server")
	srv := http.Server{
		Addr:    fmt.Sprintf(":%d", s.Port),
		Handler: s.Mux,
	}

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
