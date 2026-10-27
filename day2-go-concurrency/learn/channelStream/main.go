package channelStream

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

type Result struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

type Block struct {
	LastPrice string `json:"lastPrice,omitempty"`
	Symbol    string `json:"symbol,omitempty"`
	Price     string `json:"price,omitempty"`
	Data      []struct {
		Last string `json:"last,omitempty"`
	} `json:"data,omitempty"`
}

func BinanceQoute(ch chan float64, name string) {
	start := time.Now()
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=" + name)
	if err != nil {
		print(err)
	}
	data, err := io.ReadAll(resp.Body)
	var result Block
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		print(err)
	}
	price, err := strconv.ParseFloat(result.Price, 64)
	if err != nil {
		print(err)
	}
	ch <- price

	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)
}

func BackpackQoute(ch chan float64, name string) {
	start := time.Now()
	resp, err := http.Get("https://api.backpack.exchange/api/v1/ticker?symbol=" + name)
	if err != nil {
		print(err)
	}
	data, err := io.ReadAll(resp.Body)
	var result Block
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		print(err)
	}
	price, err := strconv.ParseFloat(result.LastPrice, 64)
	if err != nil {
		print(err)
	}
	ch <- price

	elapsed := time.Since(start)
	fmt.Println("Time taken:", elapsed)
}

func GetSolValue() (float64, error) {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=SOLUSDC")
	if err != nil {
		print(err)
	}
	data, err := io.ReadAll(resp.Body)
	var result Result
	err = json.Unmarshal([]byte(data), &result)
	if err != nil {
		print(err)
	}

	price, err := strconv.ParseFloat(result.Price, 64)
	if err != nil {
		print(err)
		return 0.0, err
	}
	return price, nil
}

func AdvancedChannel() {
	// for val := range ReturnChannel(10) {
	// 	fmt.Println(val)
	// }

	btc := make(chan float64)
	go BinanceQoute(btc, "BTCUSDC")
	sol := make(chan float64)
	go BackpackQoute(sol, "SOL_USDC")

	// get the fastest price

	select {
	case price := <-btc:
		fmt.Println("BTC price is :", price)
	case price := <-sol:
		fmt.Println("SOL price is :", price)
	case <-time.After(10 * time.Second):
		fmt.Println("Timeout")
	}
}

func ReturnChannel(n int) <-chan float64 {
	ch := make(chan float64)

	sum := 0.0
	go func() {
		for range n {
			price, err := GetSolValue()
			if err != nil {
				print(err)
			}
			fmt.Println("Price is :", price)
			sum += price
			ch <- sum
		}
		close(ch)
	}()
	return ch
}
