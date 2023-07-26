package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
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

func main() {
	// make a sample HTTP GET request
	r, err := http.Get("https://jsonplaceholder.typicode.com/users/1")
	if err != nil {
		log.Fatal(err)
	}

	// Decode Body into interface
	var v Welcome
	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(v)

	//Post
	vJSON, err := json.Marshal(v)
	b := bytes.NewReader(vJSON)
	// post some data
	r2, err := http.Post(
		"http://dummy.restapiexample.com/api/v1/create",
		"application/json",
		b,
	)
	if err != nil {
		log.Fatal(err)
	}

	// print request `Content-Type` header
	requestContentType := r2.Request.Header.Get("Content-Type")
	fmt.Println("Request content-type:", requestContentType)

	// read response data
	data, err := ioutil.ReadAll(r2.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer r2.Body.Close()

	// print response body
	fmt.Printf("%s\n", data)

	//
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("path:", r.URL.Path)
		fmt.Fprintln(w, "Hello, Client")
	}))
	defer ts.Close()
	//netClient := ts.Client()
	//q, err := http.NewRequest("GET", "http://example.com", nil)
	//response, err := netClient.Do(q)
	response, err := http.Get(ts.URL + "/test")
	if err != nil {
		log.Fatal(err)
	}
	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	fmt.Println(string(data))

}
