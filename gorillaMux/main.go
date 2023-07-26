package main

import (
	"fmt"
	"gmux/pkg/config"
	"gmux/pkg/server"
	"log"
	"net/http"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatal(err)
	}

	srv := server.NewServer(cfg)
	srv.Routes()

	if err := srv.Run(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("%v\n", err.Error())
	}

}
