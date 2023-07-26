package server

import (
	"httpTempate/pkg/helper"
	"net/http"
)

//loggerMw middleware as handler
func (s *Server) loggerMw(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uri := r.URL.String()
		method := r.Method
		//referer := r.Header.Get("Referer")
		//userAgent := r.Header.Get("User-Agent")
		//remoteAddr := r.RemoteAddr
		s.logger.Info().
			Str("Method", method).
			Str("uri", uri).
			//Str("remoteAddr", remoteAddr).
			Msg("Access")
		h.ServeHTTP(w, r)
	})
}

// mwCheck middleware has handelerfunc
func (s *Server) isAuthorized(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") == "" {
			s.logger.Warn().Str("Message", "No Token Supplied").Msg("Check Authorization")
			helper.RespondMessage(w, r, http.StatusUnauthorized, "No Token Supplied")
			return
		}
		h(w, r)
	}
}
