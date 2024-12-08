package rest

import (
	"net/http"
)

type Endpoints struct {
	Items_price string
	Items_time  string
	Categories  string
	Currencies  string
	Sellers     string
}

// Client …
type Client struct {
	BaseURL      string
	PostPipeName string
	HTTP         *http.Client
	Endpoints    Endpoints
}

func NewClient(baseURL string, endpoints Endpoints) *Client {
	return &Client{
		BaseURL:      baseURL,
		PostPipeName: "post_pipe",
		HTTP:         http.DefaultClient,
		Endpoints:    endpoints,
	}
}

func (c *Client) RestMeli_Items(id string) (*http.Response, error) {
	return nil, nil
}
