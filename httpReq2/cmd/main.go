package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"httpReq2/pkg/httpapi"
	"log"
	"net/http"
	"net/url"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Welcome template
type Welcome struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Website  string `json:"website"`
}

type Podcast struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Title  string             `json:"title,omitempty" bson:"title,omitempty"`
	Author string             `json:"author,omitempty" bson:"author,omitempty"`
}

func main() {
	getCfg := httpapi.Config{URL: "https://jsonplaceholder.typicode.com/users/1",
		Timeout:             3000,
		InsecureSkipVerify:  true,
		MaxConnsPerHost:     100,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     5000,
		MaxRetry:            5,
		RetryDelay:          100,
	}

	client, err := httpapi.New(getCfg)
	if err != nil {
		log.Fatal(err)
	}

	rel := &url.URL{Path: ""}
	resp, err := client.Get(rel, http.Header{"Accept": []string{"application/json"}})
	if err != nil {
		log.Fatal(err)
	}

	v := Welcome{}
	data := resp.Body

	err = json.Unmarshal(data, &v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	postCfg := httpapi.Config{URL: "http://dummy.restapiexample.com/api/v1/create",
		Timeout:             3000,
		InsecureSkipVerify:  true,
		MaxConnsPerHost:     100,
		MaxIdleConns:        100,
		MaxIdleConnsPerHost: 10,
		IdleConnTimeout:     5000,
		MaxRetry:            5,
		RetryDelay:          100,
	}
	client, err = httpapi.New(postCfg)
	if err != nil {
		log.Fatal(err)
	}

	resp, err = client.Post(rel, http.Header{"Content-Type": []string{"application/json"}}, bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(resp.Body))

	// mgoCfg := mgo.Config{
	// 	Database: "test",
	// 	URL:      "mongodb://localhost:27017",
	// }
	// collection := map[string]string{
	// 	"podcasts": "podcasts",
	// }

	// mongo, err := mgo.New(mgoCfg, collection)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer mongo.Disconnect()

	// podcast := Podcast{
	// 	Title:  "The Real Test",
	// 	Author: "John Smith",
	// }

	// podcastsCollection := mongo.GetCollection("podcasts")
	// insertResult, err := podcastsCollection.InsertOne(context.Background(), podcast)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(insertResult)

}
