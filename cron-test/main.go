package main

import (
	"cron-test/pkg/config"
	"cron-test/pkg/server"
	"fmt"
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
	srv.Cjob.Start()
	//srv.Cjob.AddFunc("*/5 * * * *", func() { fmt.Println("[Job 1] Every 5 sec job") })
	srv.Cjob.AddFunc("0 22 8 * * *", func() { fmt.Println("[Job 2] @ 8:22am") })

	if err := srv.Run(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("%v/n", err.Error())
	}

}
