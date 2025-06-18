package openexchangerates

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	APIKey string
}

type ClientInterface interface {
	FetchRates(currencies []string) (map[string]float64, error)
}

func NewClient(apiKey string) *Client {
	return &Client{APIKey: apiKey}
}

type RatesResponse struct {
	Rates map[string]float64 `json:"rates"`
}

func (c *Client) FetchRates(currencies []string) (map[string]float64, error) {
	url, err := buildRatesURL(c.APIKey, currencies)
	if err != nil {
		return nil, fmt.Errorf("failed to build rates url: %w", err)
	}

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch rates: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API error: status %d", resp.StatusCode)
	}

	var result RatesResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return result.Rates, nil
}
