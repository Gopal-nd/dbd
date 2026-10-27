package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
	"sync"
	"time"

	tradbrainbot "go_concurrency/fileAndFolder"
)

type Block struct {
	LastPrice string `json:"lastPrice,omitempty"`
	Symbol    string `json:"symbol,omitempty"`
	Price     string `json:"price,omitempty"`
	Result    struct {
		List []struct {
			LastPrice string `json:"lastPrice,omitempty"`
		} `json:"list,omitempty"`
		SOLUSDC struct {
			C []string `json:"c,omitempty"`
		} `json:"SOLUSDC,omitempty"`
	} `json:"result,omitempty"`
	Data []struct {
		Last string `json:"last,omitempty"`
	} `json:"data,omitempty"`
}

type ProviderData struct {
	url      string
	price    string
	provider string
}

func LoopArb(wg *sync.WaitGroup) {
	defer wg.Done()
	markets := make(map[int]ProviderData)
	markets[1] = ProviderData{url: "https://api.backpack.exchange/api/v1/ticker?symbol=SOL_USDC", price: "", provider: "backpack"}
	markets[2] = ProviderData{url: "https://api.binance.com/api/v3/ticker/price?symbol=SOLUSDC", price: "", provider: "binance"}
	markets[3] = ProviderData{url: "https://api.kraken.com/0/public/Ticker?pair=SOLUSDC", price: "", provider: "karken"}
	markets[4] = ProviderData{url: "https://api.bybit.com/v5/market/tickers?category=spot&symbol=SOLUSDC", price: "", provider: "bybit"}
	markets[5] = ProviderData{url: "https://www.okx.com/api/v5/market/ticker?instId=SOL-USDC", price: "", provider: "OKX"}

	client := &http.Client{Timeout: 10 * time.Second}

	for _, market := range markets {
		req, err := http.NewRequest("GET", market.url, nil)
		if err != nil {
			fmt.Println("error in creating a request", req)
		}
		req.Header.Add("Accept", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println(err)
		}
		bytesData, er := io.ReadAll(resp.Body)

		if er != nil {
			fmt.Println(err)
		}

		// fmt.Println(string(bytesData))
		var coins Block
		err = json.Unmarshal(bytesData, &coins)
		if err != nil {
			fmt.Println("error in unmarshaling", err)
		}

		switch market.provider {
		case "backpack":
			markets[1] = ProviderData{url: market.url, price: coins.LastPrice, provider: market.provider}
		case "binance":
			markets[2] = ProviderData{url: market.url, price: coins.Price, provider: market.provider}
		case "karken":
			markets[3] = ProviderData{url: market.url, price: coins.Result.SOLUSDC.C[0], provider: market.provider}
		case "bybit":
			markets[4] = ProviderData{url: market.url, price: coins.Result.List[0].LastPrice, provider: market.provider}
		case "OKX":
			markets[5] = ProviderData{url: market.url, price: coins.Data[0].Last, provider: market.provider}
		}
	}
	// why i am not getting the price of the market? because the price is not in the response?
	arr := []ProviderData{}

	for _, market := range markets {
		arr = append(arr, market)
	}

	minPrice := math.MaxFloat64
	maxPrice := -math.MaxFloat64
	minProvider := ""
	maxProvider := ""
	for _, market := range arr {
		p, err := strconv.ParseFloat(market.price, 64)
		if err != nil {
			fmt.Println("bad price", market.provider, err)
			continue
		}
		if p < minPrice {
			minPrice = p
			minProvider = market.provider
		}
		if p > maxPrice {
			maxPrice = p
			maxProvider = market.provider
		}
	}

	if maxPrice-minPrice >= (minPrice * 0.0025) { // 0.25% profit
		fmt.Println("Arb opportunity found")
		fmt.Println("lowest", minProvider, minPrice)
		fmt.Println("highest", maxProvider, maxPrice)
		fmt.Println("profit", maxPrice-minPrice)
		fmt.Println("profit percentage", (maxPrice-minPrice)/minPrice*100, "%")
	} else {
		fmt.Println("No arb opportunity found")
	}
}

func GetSolValue(name string, wg *sync.WaitGroup) {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=SOLUSDC")
	if err != nil {
		print(err)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		print(err)
	}
	print(string(data))

	defer wg.Done()
	fmt.Println("Hello", name)
}

func SayHello(name string) {
	fmt.Println("Hello", name)
}

func print(a any) {
	fmt.Println(a)
}

func main() {
	fmt.Println("Go Concurrency")

	var wg sync.WaitGroup

	SayHello("Gopal")
	// wg.Add(6)
	//  output order will be different
	// go GetSolValue(" go routine 1", &wg)

	// go GetSolValue(" go routine 2", &wg)

	// go GetSolValue(" go routine 3", &wg)

	// go GetSolValue(" go routine 4", &wg)

	// go GetSolValue(" go routine 5", &wg)

	// go GetSolValue(" go routine 6", &wg)
	// n := 0

	for i := range 1 {

		wg.Add(1)
		print(i)
		func() {
			start := time.Now()
			go LoopArb(&wg)
			defer fmt.Println(" - ", time.Since(start))
		}()

	}

	wg.Wait()
	SayHello("Ram")

	// channels.Chan()

	// channelStream.AdvancedChannel()

	tradbrainbot.TradBrainBot()

	fmt.Println("All Done")
	// time.Sleep(1 * time.Second)
}
