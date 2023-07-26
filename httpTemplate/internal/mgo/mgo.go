package mgo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Mgo hold mongo connection
type Mgo struct {
	Client          *mongo.Client
	DBClientTimeOut int
	DBInsertTimeOut int
	DBURI           string
	DBName          string
	DBCol           string
}

//New Returns constructed Mgo
func New() *Mgo {
	m := &Mgo{
		Client: &mongo.Client{},
	}
	return m
}

//initClient sets client to connect to config.DBURI
func (m *Mgo) initClient() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(m.DBClientTimeOut)*time.Second)
	defer cancel()
	opt := options.Client().ApplyURI(m.DBURI)
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		log.Println("unable to create mongo client connection")
		return
	}
	m.Client = client
}

//InsertDocument insert document into DB: Config.DBName; Collection: Config.DBColl
func (m *Mgo) InsertDocument(ctx context.Context, doc interface{}) (interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(m.DBInsertTimeOut)*time.Second)
	defer cancel()
	collection := m.Client.Database(m.DBName).Collection(m.DBCol)
	inserted, err := collection.InsertOne(ctx, doc)
	return inserted.InsertedID, err
}
