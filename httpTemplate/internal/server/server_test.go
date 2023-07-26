package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServer(t *testing.T) {
	srv := NewServer()
	srv.Routes()
	srv.SetupLocal()
	r := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	result := w.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("Get /, got: %d, want: %d.", result.StatusCode, http.StatusOK)
	}
}

func TestAppealHandler_fail(t *testing.T) {
	srv := NewServer()
	srv.Routes()
	srv.SetupLocal()
	r := httptest.NewRequest("GET", "/appeal", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)

	result := w.Result()
	if result.StatusCode != http.StatusUnauthorized {
		t.Errorf("Get /, got: %d, want: %d.", result.StatusCode, http.StatusUnauthorized)
	}
}
