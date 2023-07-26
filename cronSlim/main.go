package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/robfig/cron"
)

func callPost() {
	fmt.Println("Calling api")

	url := "https://post.example.com"
	contentType := "application/json"
	payload := strings.NewReader(``)

	_, err := http.Post(url, contentType, payload)

	if err != nil {
		fmt.Printf("An Error Occured %v", err)
	}
}

func main() {
	location, err := time.LoadLocation("America/Chicago")
	if err != nil {
		fmt.Println("Error loading time location. using default.")
		location = time.FixedZone("UTC", -6*3600)
	}

	fmt.Println("Start Cron")
	cr := cron.NewWithLocation(location)

	cr.AddFunc("0 5 8 * * *", callPost)
	cr.Start()

	fmt.Println("Start Server")
	http.ListenAndServe(":8090", nil)
}
