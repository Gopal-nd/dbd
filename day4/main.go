package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"sync"
	"time"
)

func fetchAll(ctx context.Context, c *http.Client, fetchers []fetcher) ([]Quote, []error) {
	ctx, cancel := context.WithTimeout(ctx, time.Duration(500*time.Millisecond))
	defer cancel()

	type Result struct {
		Quote Quote
		Error error
	}

	results := make([]Result, len(fetchers))
	var wg sync.WaitGroup

	for i, fetcher := range fetchers {
		wg.Add(1)
		go func() {
			start := time.Now()
			defer func() {
				duration := time.Since(start)
				fmt.Println(fetcher.name, "time taken", duration)
				wg.Done()
			}()
			price, err := fetcher.fn(ctx, c)
			if err != nil {
				results[i] = Result{Quote: Quote{Source: fetcher.name, Price: 0}, Error: err}
				return
			}
			results[i] = Result{Quote: Quote{Source: fetcher.name, Price: price}, Error: err}
		}()
	}
	wg.Wait()
	var qoutes []Quote
	var err []error

	for _, v := range results {
		if v.Error == nil {
			qoutes = append(qoutes, v.Quote)
		} else {
			err = append(err, v.Error)
		}
	}
	return qoutes, err
}

func main() {
	// gracefullShutdown()
	// Err()
	// fun()

	server := &http.Client{}

	fetchers := []fetcher{
		{"binance", fetchBinance},
		{"coinbase", fetchCoinbase},
		{"kraken", fetchKraken},
	}

	quotes, errs := fetchAll(context.Background(), server, fetchers)

	for _, err := range errs {
		var sourceErr *SourceError
		if errors.As(err, &sourceErr) {
			reason := ""
			if errors.Is(err, context.DeadlineExceeded) {
				reason = "Timeout"

			}
			fmt.Println("Failed to fetch from source:", sourceErr.Error(), reason)
		} else {
			fmt.Println("Error:", err)
		}
	}

	var price []float64
	for _, q := range quotes {
		fmt.Println(q.Source, q.Price)
		price = append(price, q.Price)
	}

	fmt.Println(findMedian(price))
}

func findMedian(price []float64) float64 {
	if len(price) == 0 {
		return 0.0
	}
	s := slices.Clone(price)
	fmt.Println("Spreed is :", s[len(s)-1]-s[0])
	if len(s)%2 == 0 {
		return (s[len(s)/2] + s[len(s)/2-1]) / 2
	}
	return s[len(s)/2]
}
