package helper

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//JWTRequest stuct
type JWTRequest struct {
	URL     string
	Method  string
	JWT     string
	Timeout time.Duration
}

// ExecuteJSON perform http request
func (j *JWTRequest) ExecuteJSON(v interface{}) ([]byte, error) {
	if j.URL == "" {
		return nil, errors.New("missing data: URL")
	}
	if j.Method == "" {
		j.Method = http.MethodGet
	}
	if j.JWT == "" {
		return nil, errors.New("missing data: JWT")
	}

	vJSON, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal json: %s", err)
	}

	r, err := http.NewRequest(j.Method, j.URL, bytes.NewBuffer(vJSON))
	if err != nil {
		return nil, fmt.Errorf("unable to create new request: %s", err)
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", j.JWT)

	client := &http.Client{Timeout: j.Timeout}
	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %s", err)
	}
	defer resp.Body.Close()

	return body, nil
}

// ExecuteJSON perform http request
func ExecuteJSON(url, method, jwt string, v interface{}) (*http.Response, error) {
	if url == "" {
		return nil, errors.New("missing data: URL")
	}
	if method == "" {
		method = http.MethodGet
	}
	if jwt == "" {
		return nil, errors.New("missing data: JWT")
	}

	vJSON, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("unable to marshal json: %s", err)
	}

	r, err := http.NewRequest(method, url, bytes.NewBuffer(vJSON))
	if err != nil {
		return nil, fmt.Errorf("unable to create new request: %s", err)
	}
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Authorization", jwt)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("unable to perform request: %s", err)
	}

	return resp, nil
}
