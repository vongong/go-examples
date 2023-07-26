package main

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Parent struct {
	NoteID        string `json:"noteId,omitempty" bson:"noteId,omitempty"`
	TccInstanceID string `json:"tccInstanceId" bson:"tccInstanceId"`
}
type Podcast struct {
	Parent `bson:",inline"`
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Author string             `json:"author,omitempty" bson:"author,omitempty"`
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	opt := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, opt)
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)
	database := client.Database("test")
	podcastsCollection := database.Collection("podcasts")

	podcast := Podcast{
		Title:  "The Real Test",
		Author: "John Smith",
	}
	podcast.NoteID = "123456"
	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult)

}
