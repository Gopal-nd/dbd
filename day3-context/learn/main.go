package main

import (
	"context"
	"fmt"
	"time"
)

// enums

type Side uint8

const (
	Buy Side = iota
	Sell
	Hold
)

func (s Side) String() string {
	switch s {
	case Buy:
		return "BUY"
	case Sell:
		return "SELL"
	case Hold:
		return "HOLD"
	default:
		return "UNKNOWN"
	}
}

func GetPricefeed(ctx context.Context, prices chan int) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	price := 100

	for {
		select {
		case <-ctx.Done():
			fmt.Println("feed stoped", ctx.Err())
			return

		case <-ticker.C:
			price++
			select {
			case prices <- price:
			case <-ctx.Done():
				fmt.Println("feed stoped", ctx.Err())
				return

			}

		}
	}
}

func SlowApiCall(ctx context.Context) (error, int) {
	select {
	case <-ctx.Done():
		err := fmt.Errorf(" %w", ctx.Err())
		return err, 0
	case <-time.After(1 * time.Second):
		return nil, 100
	}
}

func main() {
	// s := Side(10)
	// fmt.Println(s.String())

	// ctx, cancel := context.WithCancel(context.Background())
	// prices := make(chan int)
	// go GetPricefeed(ctx, prices)

	ctx, cancel := context.WithTimeout(context.Background(), 1500*time.Millisecond)
	err, result := SlowApiCall(ctx)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("result: ", result)
	defer cancel()
	cancel()
	time.Sleep(3 * time.Second)
}
