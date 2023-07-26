package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

const (
	//PortEnvVar is Exported value
	PortEnvVar = "ExportedValue"
)
const (
	portDefault    = "8080"
	versionDefault = "0.1.0"
	appDefault     = "Template"
)

// Config hold config data
type Config struct {
	Application string
	Version     string
	ApiTNG      string
	Port        string
	store       *viper.Viper
}

//NewConfig things
func NewConfig() *Config {
	v := viper.New()
	v.SetDefault("PORT", portDefault)
	v.SetDefault("VERSION", versionDefault)
	v.SetDefault("APP", appDefault)
	v.SetDefault("ApiTNG", "000")
	cfg := &Config{
		Application: appDefault,
		Version:     versionDefault,
		Port:        portDefault,
		store:       v,
	}
	return cfg
}

func (c *Config) refresh() {
	c.Application = c.store.GetString("APP")
	c.Port = c.store.GetString("PORT")
	c.Version = c.store.GetString("VERSION")
	c.ApiTNG = c.store.GetString("APITNG")
}

//FromEnv load config from Env
func (c *Config) FromEnv() {
	//simulate enviroment var set.
	os.Setenv("APITNG", "https://tngs-authorization.int-capi-ranchers.cncpl.us/v2")
	c.store.AutomaticEnv()
	c.refresh()
}

//FromJSONFile load from file
func (c *Config) FromJSONFile(filename string) error {
	c.store.SetConfigFile(filename)
	c.store.SetConfigType("json")
	err := c.store.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("Could not file file: config.json")
		}
		return fmt.Errorf("unable to read file: %s", err)
	}
	c.refresh()
	return nil
}
