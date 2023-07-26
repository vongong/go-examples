package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type vertex struct {
	Lat, Long float64
}
type vertMap map[string]*vertex

////////////////////////
type routeKey struct {
	Method string
	URL    string
}

type responseMap map[routeKey]*http.Response

//MockClient Hold infor for mock client
type MockClient struct {
	DefaultResponse *http.Response
	ResponseMap     responseMap
	mu              sync.RWMutex
}

//NewMockClient return MockClient initialized
func NewMockClient() *MockClient {
	c := &MockClient{
		ResponseMap:     make(responseMap),
		DefaultResponse: &http.Response{},
	}
	return c
}
func (c *MockClient) getKey(r *http.Request) routeKey {
	key := routeKey{
		Method: strings.ToUpper(r.Method),
		URL:    r.URL.String(),
	}
	return key

}

//Do looks up request and returns reponse, if not found return default
func (c *MockClient) Do(r *http.Request) (*http.Response, error) {
	key := c.getKey(r)
	c.mu.RLock()
	defer c.mu.RUnlock()
	response, found := c.ResponseMap[key]
	if found {
		return response, nil
	}

	return c.DefaultResponse, nil
}

//SetDefaultResponse set default response
func (c *MockClient) SetDefaultResponse(response *http.Response) {
	c.DefaultResponse = response
}

//RegisterResponse Register response with request
func (c *MockClient) RegisterResponse(r *http.Request, response *http.Response) {
	key := c.getKey(r)
	c.mu.Lock()
	c.ResponseMap[key] = response
	c.mu.Unlock()
}

//NewResponse returns http reponse
func NewResponse(status int, data []byte) *http.Response {
	response := &http.Response{
		Status:        strconv.Itoa(status),
		StatusCode:    status,
		Body:          ioutil.NopCloser(bytes.NewReader(data)),
		Header:        http.Header{},
		ContentLength: -1,
	}
	return response
}

func main() {
	m := make(vertMap)

	m["Bell Labs"] = &vertex{100.10, 22.90}
	m["Bell Labs"].Lat++
	fmt.Println(*m["Bell Labs"])

	v, found := m["Bell Labs"]
	fmt.Println(*v)
	fmt.Println(found)

}
