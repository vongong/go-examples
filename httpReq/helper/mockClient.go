package helper

import (
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//MockClient stuct of MockClient
type MockClient struct {
	DefaultResponse *http.Response
	Timeout         time.Duration
}

//NewMockClient return MockClient Initialized
func NewMockClient() *MockClient {
	c := &MockClient{
		DefaultResponse: &http.Response{},
	}
	return c
}

//Do returns default Response
func (c *MockClient) Do(req *http.Request) (*http.Response, error) {
	//return &http.Response{}, nil
	return c.DefaultResponse, nil
}

//Get calls Do with Get Method
func (c *MockClient) Get(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

//Head calls Do with Head Method
func (c *MockClient) Head(url string) (resp *http.Response, err error) {
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

//Post calls Do with Post Method
func (c *MockClient) Post(url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", contentType)
	return c.Do(req)
}

//PostForm calls Post with data encoded
func (c *MockClient) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	return c.Post(url, "application/x-www-form-urlencoded", strings.NewReader(data.Encode()))
}
