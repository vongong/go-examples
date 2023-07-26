package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
)

type config struct {
	Port int
	Name string
}

func main() {
	v := viper.New()
	v.SetDefault("Port", 12346)

	var cfg config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Println(err)
	}

	fmt.Println(cfg)
}
