package main

import (
	"fmt"
	"httpTempate/internal/server"
	"net/http"
)

func main() {
	srv := server.NewServer()
	srv.Routes()
	srv.SetupLocal()

	if err := srv.Start(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("%v/n", err.Error())
	}

	//curl localhost:8080
	//curl localhost:8080/hello
	//curl localhost:8080/appeal
	//curl localhost:8080/appeal --header 'Authorization: 123'
}
