package mssql

import (
	"context"
	"database/sql"
	"errors"
	"net/url"
	"time"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/rs/zerolog/log"
)

// Config represents the basic configurables for the sql connection.
type Config struct {
	Database string  `json:"database"`
	URL      string  `json:"url" base64:"true"`
	Username string  `json:"username" base64:"true"`
	Password string  `json:"password" base64:"true"`
	Options  Options `json:"options"`
}

// Options represents the configurable sql database connection pool options.
type Options struct {
	MaxOpenConnections int           `json:"max-open-connections"`
	MaxIdleConnections int           `json:"max-idle-connections"`
	MaxLifetime        time.Duration `json:"max-lifetime-ms"`
}

type Sql struct {
	name     string
	host     string
	username string

	URL      *url.URL
	Options  Options
	Database *sql.DB // Export for mocking
}

// Connect creates a new sql connection configurables.
func New(conf Config) (*Sql, error) {
	db := Sql{}
	uri, err := url.Parse(conf.URL)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Parse URL")
		return nil, err
	}

	db.URL = uri
	db.Options = Options{
		MaxOpenConnections: conf.Options.MaxOpenConnections,
		MaxIdleConnections: conf.Options.MaxIdleConnections,
		MaxLifetime:        conf.Options.MaxLifetime * time.Millisecond,
	}

	sqlDriver := db.URL.Scheme

	database, err := sql.Open(sqlDriver, db.URL.String())
	if err != nil {
		log.Error().Err(err).Msg("Unable to Open DB Connection")
		return nil, err
	}

	database.SetMaxOpenConns(db.Options.MaxOpenConnections)
	database.SetMaxIdleConns(db.Options.MaxIdleConnections)
	database.SetConnMaxLifetime(db.Options.MaxLifetime)
	db.Database = database
	db.name, db.host, db.username = db.URL.Query().Get("database"), db.URL.Hostname(), db.URL.User.Username()

	return &db, nil
}

// Ping is used to check if the remote server is available.
func (ds *Sql) Ping() error {
	if ds.Database == nil {
		err := errors.New("could not connect to database")
		log.Error().Err(err)
		return err
	}

	// creates a new context to add to the client's connection. We do not need to return a CancelFunc().
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err := ds.Database.PingContext(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Ping")
		return err
	}
	return nil
}

// QueryContext wraps and exposes sql.QueryContext with metrics
// func (ds *Sql) QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error) {
// 	return ds.Database.QueryContext(ctx, query, args...)
// }
