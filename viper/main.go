package main

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

func main() {
	os.Setenv("SPF_ID", "13") // typically done outside of the app
	fmt.Println("os getenv:", os.Getenv("SPF_ID"))

	viper.SetDefault("SPF_ID", "42")
	fmt.Println("viper SPF_ID get:", viper.GetString("SPF_ID"))
	fmt.Println("viper Port get:", viper.GetString("PORT"))
	fmt.Println("bind env")
	viper.AutomaticEnv()
	fmt.Println("viper SPF_ID get:", viper.GetString("SPF_ID"))
	fmt.Println("viper Port get:", viper.GetString("PORT"))

	viper.SetConfigFile("config.json")
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Could not file file: config.json")
		} else {
			fmt.Printf("couldn't read file: %s", err.Error())
		}
	}
	fmt.Println("viper SPF_ID get:", viper.GetString("SPF_ID"))
	fmt.Println("viper Port get:", viper.GetString("PORT"))

}
