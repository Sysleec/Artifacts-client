package artsapi

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog/log"
	"io"
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

func (c *Client) PostReq(path string, body []byte) ([]byte, error) {
	url := baseURL + path

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Err(err).Msg("failed to close response body")
		}
	}(res.Body)

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Error().Msgf("response: %s", string(resp)) // for debugging purposes
		return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}

	return resp, nil
}

func (c *Client) GetReq(path string) ([]byte, error) {
	url := baseURL + path

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.token))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %w", err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Err(err).Msg("failed to close response body")
		}
	}(res.Body)

	resp, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		log.Error().Msgf("response: %s", string(resp)) // for debugging purposes
		return nil, fmt.Errorf("wrong status code: %d", res.StatusCode)
	}

	return resp, nil
}
