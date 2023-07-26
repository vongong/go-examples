package sql

import "time"

type Config struct {
	Database string  `json:"database"`
	URL      string  `json:"url" base64:"true"`
	Username string  `json:"username" base64:"true"`
	Password string  `json:"password" base64:"true"`
	Options  Options `json:"options"`
}

type Options struct {
	MaxOpenConnections int           `json:"max-open-connections"`
	MaxIdleConnections int           `json:"max-idle-connections"`
	MaxLifetime        time.Duration `json:"max-lifetime-ms"`
}
