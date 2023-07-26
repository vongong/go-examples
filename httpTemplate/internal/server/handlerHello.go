package server

import (
	"httpTempate/pkg/helper"
	"net/http"
)

func (s *Server) helloHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slogger := s.logger.With().Str("fn", "helloHandler").Logger()
		slogger.Info().Msg("hello")
		slogger.Warn().Msg("this is a warning")
		helper.RespondMessage(w, r, http.StatusOK, "Hello World from func")
	}
}
