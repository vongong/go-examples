package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"httpTempate/internal/server/config"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
)

const timeoutHTTP = 10

//Server hold relavant info
type Server struct {
	//db     *someDatabase
	router    *mux.Router
	netClient *http.Client
	logger    *zerolog.Logger
	config    *config.Config
}

//NewServer returns Server
func NewServer() *Server {
	svr := &Server{
		router:    mux.NewRouter(),
		netClient: &http.Client{Timeout: timeoutHTTP * time.Second},
		config:    config.NewConfig(),
	}
	//svr.Routes()
	return svr
}

//SetupLocal setups server for localhost
func (s *Server) SetupLocal() {
	//Setup Logger
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	lgr := zerolog.New(output).With().Timestamp().Logger()
	s.logger = &lgr
	//Setup Config
	s.config.FromEnv()
	//Setup Database
	//Setup Router
}

//Routes routes for server
func (s *Server) Routes() {
	s.router.HandleFunc("/", s.helloHandler())
	s.router.Handle("/hello", s.helloHandler())
	s.router.HandleFunc("/appeal", s.isAuthorized(s.appealHandler()))
	s.router.Use(s.loggerMw)
}

//Start starts server
func (s *Server) Start() error {
	s.logger.Info().Str("port", s.config.Port).Msg("Starting Server")

	srv := &http.Server{
		Addr:    ":" + s.config.Port,
		Handler: s.router,
	}

	//Graceful shutdown
	go func() {
		done := make(chan os.Signal, 1)
		signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		<-done
		s.logger.Info().Msg("Shutting Down Server")

		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("HTTP server Shutdown: %v", err)
		}
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		return fmt.Errorf("HTTP server ListenAndServe: %v", err)
	}
	return nil
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
