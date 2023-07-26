package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoFields struct {
	Key string `json:"key,omitempty"`
	// ObjectId() or objectid. ObjectID is deprecated--use primitive instead
	ID int `bson:"_id, omitempty"`

	// Use these field tags so Golang knows how to map MongoDB fields
	// `bson:"string field" json:"string field"`
	Name  string `bson:"name" json:"name"`
	Color string `bson:"color" json:"color"`
	Qty   int    `bson:"qty" json:"qry"`
}

func main() {

	// Create a string using ` string escape ticks
	query := `{"name":{"$eq":"item1"}}`

	// Declare an empty BSON Map object
	var filter bson.M

	// Use the JSON package's Unmarshal() method
	err := json.Unmarshal([]byte(query), &filter)
	if err != nil {
		log.Fatal("json. Unmarshal() ERROR:", err)
	} else {
		fmt.Println("bsonMap:", filter)
		fmt.Println("bsonMap TYPE:", reflect.TypeOf(filter))
		fmt.Println("BSON:", reflect.TypeOf(bson.M{"int field": bson.M{"$gt": 42}}))
	}

	// Declare host and port options to pass to the Connect() method
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to the MongoDB and return Client instance
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		fmt.Println("mongo. Connect() ERROR:", err)
		log.Fatal(err)
	}

	// Declare Context object for the MongoDB API calls
	ctx := context.Background()

	// Access a MongoDB collection through a database
	col := client.Database("test").Collection("foo")

	cursor, err := col.Find(ctx, filter)
	if err != nil {
		log.Fatal("col. Find ERROR:", err)
	}

	// iterate through all documents
	for cursor.Next(ctx) {
		var p MongoFields

		// Decode the document
		if err := cursor.Decode(&p); err != nil {
			log.Fatal("cursor. Decode ERROR:", err)
		}
		// Print the results of the iterated cursor object
		fmt.Printf("\nMongoFields: %+v\n", p)
	}

	endDt := time.Now()
	exists := bson.D{{Key: "authorizationNbr", Value: bson.D{{Key: "$exists", Value: true}}}}
	createdTsLTE := bson.D{{Key: "createdTs", Value: bson.D{{Key: "$lte", Value: endDt}}}}
	andList := bson.A{exists}
	andList = append(andList, createdTsLTE)
	fmt.Println("andList: ", andList)
}
