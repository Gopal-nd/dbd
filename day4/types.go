package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Quote struct {
	Source string
	Price  float64
}

type fetcher struct {
	name string
	fn   func(ctx context.Context, c *http.Client) (float64, error)
}

func getJSON(ctx context.Context, c *http.Client, url string, data any) error {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return err
	}

	resp, err := c.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(data)
}
