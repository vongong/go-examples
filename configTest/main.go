package main

import (
	"configtest/pkg/config"
	"fmt"
	"log"
)

func main() {
	cfg, err := config.GetConfig()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
