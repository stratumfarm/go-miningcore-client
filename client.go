package miningcore

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// ClientOpts are options for the client.
type ClientOpts func(*Client)

// WithoutTLSVerify disables TLS verification.
func WithoutTLSVerfiy() ClientOpts {
	return func(c *Client) {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}
		c.http.Transport = tr
	}
}

// WithTimouts sets the default request timeout
func WithTimeout(t time.Duration) ClientOpts {
	return func(c *Client) {
		c.timeout = t
	}
}

// Client represents a client for the miningcore API.
type Client struct {
	timeout time.Duration
	url     string
	http    *http.Client
}

// New create a new client for the miningcore API.
func New(url string, opts ...ClientOpts) *Client {
	c := &Client{
		timeout: time.Second * 20,
		url:     url,
	}
	for _, opt := range opts {
		opt(c)
	}
	c.http = &http.Client{Timeout: c.timeout}
	return c
}

// doRequest performs the actual request to the miningcore API.
func (c *Client) doRequest(ctx context.Context, endpoint, method string, expRes, reqData interface{}) (int, error) {
	callURL := fmt.Sprintf("%s%s", c.url, endpoint)

	var dataReq []byte
	var err error

	if reqData != nil {
		dataReq, err = json.Marshal(reqData)
		if err != nil {
			return 0, err
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, callURL, bytes.NewBuffer(dataReq))
	if err != nil {
		return 0, err
	}

	resp, err := c.http.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	switch resp.StatusCode {
	case 200:
		if expRes != nil {
			err = json.Unmarshal(body, expRes)
			if err != nil {
				return 0, err
			}
		}
		return resp.StatusCode, nil

	default:
		return resp.StatusCode, fmt.Errorf("%s", body)
	}
}
