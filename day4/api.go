package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// {"symbol":"BTCUSDT","price":"68123.45000000"}
func fetchBinance(ctx context.Context, c *http.Client) (float64, error) {

	var body struct {
		Price string `json:"price"` // a STRING, not a number
	}
	if err := getJSON(ctx, c, "https://api.binance.com/api/v3/ticker/price?symbol=BTCUSDT", &body); err != nil {
		return 0, err
	}
	return strconv.ParseFloat(body.Price, 64)
}

// {"data":{"amount":"68100.12","base":"BTC","currency":"USD"}}
func fetchCoinbase(ctx context.Context, c *http.Client) (float64, error) {
	var body struct {
		Data struct {
			Amount string `json:"amount"`
		} `json:"data"`
	}
	if err := getJSON(ctx, c, "https://api.coinbase.com/v2/prices/BTC-USD/spot", &body); err != nil {
		return 0, err
	}
	return strconv.ParseFloat(body.Data.Amount, 64)
}

// {"error":[],"result":{"XXBTZUSD":{"c":["68110.1","0.002"], ...}}}
func fetchKraken(ctx context.Context, c *http.Client) (float64, error) {
	var body struct {
		Error  []string            `json:"error"`
		Result map[string]struct { // keyed by Kraken's odd pair name
			C []string `json:"c"` // c = last trade: [price, lot volume]
		} `json:"result"`
	}
	if err := getJSON(ctx, c, "https://api.kraken.com/0/public/Ticker?pair=XBTUSD", &body); err != nil {
		return 0, err
	}
	// Kraken returns HTTP 200 even on errors; they're in the body
	if len(body.Error) > 0 {
		return 0, fmt.Errorf("api error: %s", strings.Join(body.Error, ", "))
	}
	for _, t := range body.Result { // only one entry; we don't hardcode "XXBTZUSD"
		if len(t.C) > 0 {
			return strconv.ParseFloat(t.C[0], 64)
		}
	}
	return 0, errors.New("price missing in response")
}
