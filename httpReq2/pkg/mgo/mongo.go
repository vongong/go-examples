package mgo

import (
	"context"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Config for Mongo
type Config struct {
	Database string `json:"database"`
	URL      string `json:"url" base64:"true"`
	Username string `json:"username" base64:"true"`
	Password string `json:"password" base64:"true"`
}

//Mongo Model
type Mongo struct {
	name        string
	host        string
	username    string
	URL         *url.URL
	Database    *mongo.Database
	collections map[string]*mongo.Collection
}

//New ...
func New(conf Config, collections map[string]string) (*Mongo, error) {
	db := Mongo{collections: make(map[string]*mongo.Collection)}
	if err := db.Connect(conf); err != nil {
		return nil, err
	}

	for key, collection := range collections {
		db.collections[key] = db.Database.Collection(collection)
	}
	return &db, nil
}

//Connect ...
func (db *Mongo) Connect(conf Config) error {
	uri, err := url.Parse(conf.URL)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Parse MongoDB URL")
		return err
	}

	db.URL = uri
	opts := options.Client()
	opts.ApplyURI(db.URL.String())

	client, err := mongo.NewClient(opts)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Create MongoDB Client")
		return err
	}

	// creates a new context to add to the client's connection. We do not need to return a CancelFunc().
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Error().Err(err).Msg("Unable to Connect to MongoDB")
		return err
	}

	db.Database = client.Database(conf.Database)
	db.name = db.Database.Name()
	db.host = db.URL.Hostname()
	db.username = db.URL.User.Username()

	return nil
}

// Disconnect mongodb connection
func (db *Mongo) Disconnect() {
	db.Database.Client().Disconnect(context.Background())
}

// GetCollection ...
func (db *Mongo) GetCollection(key string) *mongo.Collection {
	return db.collections[key]
}
