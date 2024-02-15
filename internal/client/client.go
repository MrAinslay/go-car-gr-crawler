package client

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
}

func New(timeout time.Duration) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}
