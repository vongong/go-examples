package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path"
	"time"
)

func main() {
	var (
		BaseUrl string
		AuthId  string
		data    []byte
	)
	u, err := url.Parse(BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	if AuthId != "" {
		u.Path = path.Join(u.Path, "authorization", AuthId)
	}

	r, err := http.NewRequest(http.MethodPost, u.String(), bytes.NewBuffer(data))
	if err != nil {
		panic(err)
	}
	r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	r.Header.Set("Content-Type", "application/json")

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	resp, err := client.Do(r)
	if err != nil {
		panic(err)
	}
	// return resp, nil
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", string(body))

	//var u user
	//err = json.NewDecoder(resp.Body).Decode(&u)
	//defer resp.Body.Close
}
