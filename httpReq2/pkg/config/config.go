package config

import (
	"encoding/json"
	"errors"
	"os"
	"reflect"

	"github.com/rs/zerolog/log"
)

//Config ...
type Config struct {
	Name     string `json:"name"`
	Port     int    `json:"port"`
	Debug    bool   `json:"debug"`
	LogLevel string `json:"log-level"`
}

// GetConfig mappings for app
func GetConfig() (Config, error) {
	conf := Config{}
	if err := read("config/config.json", &conf); err != nil {
		return conf, err
	}
	return conf, nil
}

// Read opens and reads a file into config struct based on path location.
func read(path string, conf interface{}) error {
	if reflect.ValueOf(conf).Kind() != reflect.Ptr {
		return errors.New("interface is not a pointer")
	}

	file, err := os.Open(path)
	if err != nil {
		log.Error().Stack().Caller().Err(err).Msgf("unable to open file: %s", path)
		return err
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(conf)
	if err != nil {
		log.Error().Stack().Caller().Err(err).Msgf("unable to parse file: %s", path)
		return err
	}

	return nil
}
