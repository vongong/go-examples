package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
)

func main() {
	req, err := http.NewRequest("GET", "http://api.example.com", nil)
	if err != nil {
		log.Print(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("api_key", "12345")
	q.Add("another_thing", "foo & bar")
	req.URL.RawQuery = q.Encode()

	fmt.Println(req.URL.String())

	u1, err := url.Parse("https://wfm.org/?a=1&b=2&c=3")
	if err != nil {
		log.Fatal(err)
	}
	q1 := u1.Query()
	u2, err := url.Parse("https://example.org/")
	if err != nil {
		log.Fatal(err)
	}
	u2.Path = path.Join(u2.Path, "groups")
	u2.RawQuery = q1.Encode()
	fmt.Println(u2)

}
