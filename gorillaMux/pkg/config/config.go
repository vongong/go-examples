package config

import (
	"encoding/json"
	"os"
)

type Config struct {
	Name string `json:"name"`
	Port int    `json:"port"`
}

// GetConfig mappings for app
func GetConfig() (Config, error) {
	conf := Config{}

	file, err := os.Open("config/config.json")
	if err != nil {
		return conf, err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		return conf, err
	}

	return conf, nil
}
