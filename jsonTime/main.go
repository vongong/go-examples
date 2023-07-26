package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type response1 struct {
	Assignee  string     `json:"assignee"`
	CreatedDt *time.Time `json:"createdDt"`
}

func main() {
	jsonData := []byte(`{
		"assignee": "Bob",
		"createdDt": "2021-11-16T16:42:41.394Z"
	}`)
	var resp response1
	err := json.Unmarshal(jsonData, &resp)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Assignee)
	fmt.Println(resp.CreatedDt)

	var resp2 response1
	jsonData2 := []byte(`{
		"assignee": "Alice"
	}`)

	err = json.Unmarshal(jsonData2, &resp2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp2.Assignee)
	fmt.Println(resp2.CreatedDt)

	var startTime *time.Time

	fmt.Println("Hello, playground")
	t, err := time.Parse(time.RFC3339, "2021-11-16T16:42:41.394Z")

	if err != nil {
		fmt.Println(err)
	} else {
		startTime = &t
	}

	fmt.Println(startTime)
}
