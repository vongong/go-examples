package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAppealHandler2(t *testing.T) {
	srv := NewServer()
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, client")
	}))
	defer ts.Close()
	srv.netClient = ts.Client()

	//skip middleware
	srv.router.HandleFunc("/appeal", srv.appealHandler())
	srv.SetupLocal()

	r := httptest.NewRequest("GET", "/appeal", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)

	result := w.Result()
	if result.StatusCode != http.StatusOK {
		t.Errorf("Get /, got: %d, want: %d.", result.StatusCode, http.StatusOK)
	}

}
