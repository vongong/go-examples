package httpapi

import (
	"context"
	"crypto/tls"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/rs/zerolog/log"
)

//Response Model
type Response struct {
	Status           string // e.g. "200 OK"
	StatusCode       int    // e.g. 200
	Proto            string // e.g. "HTTP/1.0"
	ProtoMajor       int    // e.g. 1
	ProtoMinor       int    // e.g. 0
	Header           http.Header
	Body             []byte
	ContentLength    int64
	TransferEncoding []string
	Uncompressed     bool
	Trailer          http.Header
	Duration         time.Duration
	Request          *http.Request
}

//Client Model
type Client struct {
	client     *http.Client
	url        *url.URL
	headers    http.Header
	maxRetry   int
	retryDelay time.Duration
}

//Config Model
type Config struct {
	URL string `json:"url"`

	// Timeout specifies a time limit for requests made by this
	// Client. The timeout includes connection time, any
	// redirects, and reading the response body. The timer remains
	// running after Get, Head, Post, or Do return and will
	// interrupt reading of the Response.Body.
	// A Timeout of zero means no timeout.
	Timeout time.Duration `json:"timeout-ms"`

	// IdleConnTimeout is the maximum amount of time an idle
	// (keep-alive) connection will remain idle before closing
	// itself.
	// Zero means no limit.
	IdleConnTimeout time.Duration `json:"idle-connection-timeout-ms"`

	// InsecureSkipVerify controls whether a client verifies the
	// server's certificate chain and host name.
	// If InsecureSkipVerify is true, TLS accepts any certificate
	// presented by the server and any host name in that certificate.
	InsecureSkipVerify bool `json:"insecure-skip-verify"`

	// MaxConnsPerHost optionally limits the total number of
	// connections per host, including connections in the dialing,
	// active, and idle states. On limit violation, dials will block.
	// Zero means no limit.
	MaxConnsPerHost int `json:"max-connection-per-host"`

	// MaxIdleConns controls the maximum number of idle (keep-alive)
	// connections across all hosts. Zero means no limit.
	MaxIdleConns int `json:"max-idle-connections"`

	// MaxIdleConnsPerHost, if non-zero, controls the maximum idle
	// (keep-alive) connections to keep per-host. If zero,
	// DefaultMaxIdleConnsPerHost is used.
	MaxIdleConnsPerHost int `json:"max-idle-connections-per-host"`

	Headers    map[string]string `json:"default-headers"`
	MaxRetry   int               `json:"max-retry"`
	RetryDelay time.Duration     `json:"retry-delay-ms"`
}

//New ...
func New(conf Config) (*Client, error) {
	uri, err := url.Parse(conf.URL)
	if err != nil {
		return nil, err
	}

	// Transport specifies the mechanism by which individual HTTP requests are made
	transport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: conf.InsecureSkipVerify,
		},
		MaxConnsPerHost:     conf.MaxConnsPerHost,
		MaxIdleConns:        conf.MaxIdleConns,
		MaxIdleConnsPerHost: conf.MaxIdleConnsPerHost,
		IdleConnTimeout:     conf.IdleConnTimeout * time.Millisecond,
	}

	// Http client
	client := &http.Client{
		Transport: transport,
		Timeout:   conf.Timeout * time.Millisecond,
	}

	// A Header represents the key-value pairs in an HTTP header. We are creating all string key pairs headers that are being read in from the config file.
	headers := http.Header{}
	for k, v := range conf.Headers {
		headers.Set(k, v)
	}

	return &Client{client: client, url: uri, headers: headers, maxRetry: conf.MaxRetry, retryDelay: conf.RetryDelay}, nil
}

// Get helper method for making a GET request
func (c *Client) Get(rel *url.URL, headers http.Header) (*Response, error) {
	return c.Do(context.Background(), http.MethodGet, rel, headers, nil)
}

// GetWithContext helper method for making a GET request
func (c *Client) GetWithContext(ctx context.Context, rel *url.URL, headers http.Header) (*Response, error) {
	return c.Do(ctx, http.MethodGet, rel, headers, nil)
}

// Post helper method for making a POST request
func (c *Client) Post(rel *url.URL, headers http.Header, body io.Reader) (*Response, error) {
	return c.Do(context.Background(), http.MethodPost, rel, headers, body)
}

// PostWithContext helper method for making a POST request
func (c *Client) PostWithContext(ctx context.Context, rel *url.URL, headers http.Header, body io.Reader) (*Response, error) {
	return c.Do(ctx, http.MethodPost, rel, headers, body)
}

// Do ...
func (c *Client) Do(ctx context.Context, method string, rel *url.URL, headers http.Header, body io.Reader) (*Response, error) {
	uri := c.url.ResolveReference(rel)

	request, err := http.NewRequest(method, uri.String(), body)
	if err != nil {
		return nil, err
	}

	request = request.WithContext(ctx)

	if headers == nil {
		headers = http.Header{}
	}

	for k, vs := range c.headers {
		for _, v := range vs {
			headers.Add(k, v)
		}
	}
	request.Header = headers
	return c.do(request)
}

func (c *Client) do(request *http.Request) (*Response, error) {
	defer func() {
		if request.Body != nil {
			request.Body.Close()
		}
	}()

	response, err := c.handle(request)
	if (err != nil || IsServerError(response.StatusCode)) && c.maxRetry > 0 {
		if err == context.Canceled {
			return response, err
		}
		log.Error().Err(err).Str("url", request.URL.String()).Msg("Error with request. Retrying...")
		retries := 0
		for retries < c.maxRetry {
			time.Sleep(c.retryDelay)

			retries++
			response, err = c.handle(request)
			if err != nil {
				continue
			}

			if !IsServerError(response.StatusCode) {
				return response, err
			}
		}
		//Max retries reached
		log.Error().Err(err).Str("url", request.URL.String()).Int("retries", retries).Msg("Max retries reached")
	}

	return response, err
}

func (c *Client) handle(request *http.Request) (*Response, error) {
	start := time.Now()
	req := request.Clone(request.Context())
	if request.Body != nil && request.GetBody != nil {
		if body, err := request.GetBody(); err == nil {
			req.Body = ioutil.NopCloser(body)
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		if err, ok := err.(*url.Error); ok {
			return nil, err.Unwrap()
		}
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := Response{
		Status:           resp.Status,
		StatusCode:       resp.StatusCode,
		Proto:            resp.Proto,
		ProtoMajor:       resp.ProtoMajor,
		ProtoMinor:       resp.ProtoMinor,
		Header:           resp.Header,
		Body:             body,
		ContentLength:    resp.ContentLength,
		TransferEncoding: resp.TransferEncoding,
		Uncompressed:     resp.Uncompressed,
		Trailer:          resp.Trailer,
		Duration:         time.Since(start),
		Request:          request,
	}
	return &response, nil

}

//inRange checks if  a <= code < b
func inRange(code, a, b int) bool {
	return a <= code && code < b
}

// IsSuccessful checks if server code being passed is a successfully code
func IsSuccessful(code int) bool {
	return inRange(code, 200, 300)
}

// IsServerError checks if server code being passed is a server error code
func IsServerError(code int) bool {
	return inRange(code, 500, 600)
}

// IsClientError checks if server code being passed is a client error code
func IsClientError(code int) bool {
	return inRange(code, 400, 500)
}
