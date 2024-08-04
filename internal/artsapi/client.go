package artsapi

import (
	"net/http"
	"time"
)

type Client struct {
	httpClient http.Client
	token      string
}

func NewClient(interval time.Duration, token string) Client {
	return Client{
		httpClient: http.Client{
			Timeout: interval,
		},
		token: token,
	}
}
