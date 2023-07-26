package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting server")

	//v1
	//http.HandleFunc("/", handler)
	//http.ListenAndServe(":8080", nil)

	//v2
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)
	mux.HandleFunc("/auth", mwCheck(authHandler))
	srv := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
	//Client - curl localhost:8080
}

func mwCheck(f func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("mw Check")
		f(w, r)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server: hello handler started")
	fmt.Fprintln(w, "Hello World")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("server: auth handler started")
	fmt.Fprintln(w, "auth")
}
