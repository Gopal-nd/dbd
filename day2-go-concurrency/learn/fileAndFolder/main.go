package tradbrainbot

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Result struct {
	Symbol   string `json:"symbol"`
	Price    string `json:"price"`
	Time     string `json:"time"`
	Provider string `json:"provider"`
}

type Block struct {
	LastPrice string `json:"lastPrice,omitempty"`
	Symbol    string `json:"symbol,omitempty"`
	Price     string `json:"price,omitempty"`
	Data      []struct {
		Last string `json:"last,omitempty"`
	} `json:"data,omitempty"`
}

func BinanceQoute(ch chan Result, name string, wg *sync.WaitGroup) {
	defer wg.Done()
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
	ch <- Result{Symbol: name, Price: strconv.FormatFloat(price, 'f', -1, 64), Time: start.Format(time.RFC3339), Provider: "Binance"}

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

func createDir() {
	err := os.Mkdir("data", 0o755)
	if err != nil {
		fmt.Println("Error creating directory:", err)
	}
}

func createFile(data ...string) string {
	var ext string
	if data == nil {
		ext = ".txt"
	} else {
		ext = ".csv"
	}

	name := "data/" + time.Now().Format("2006-01-02_15-04-05") + strings.ToLower(ext)
	file, _ := os.Create(name)

	for _, d := range data {
		_, _ = file.WriteString(d + "\n")
	}

	defer file.Close()

	return name
}

// func createNestedDir() {
// 	err := os.MkdirAll("data/2026/10/27", 0o755)
// 	if err != nil {
// 		fmt.Println("Error creating nested directory:", err)
// 	}
// }

func appendToFile(name string, data ...string) *os.File {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}

	for _, d := range data {
		_, _ = file.WriteString(d + "\n")
	}

	return file
}

func readFromFile(name string) string {
	file, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	return string(file)
}

func GenerateCSV() {
	name := "data/" + time.Now().Format("2006-01-02_15-04-05") + ".csv"

	file, err := os.Create(name)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	records := [][]string{
		{"Symbol", "Price", "Time", "Provider"},
	}

	tokens := make(chan Result, 100)
	var wg sync.WaitGroup
	for range 12 {
		wg.Add(12)
		go BinanceQoute(tokens, "BTCUSDC", &wg)
		go BinanceQoute(tokens, "SOLUSDC", &wg)
		go BinanceQoute(tokens, "ETHUSDC", &wg)
		go BinanceQoute(tokens, "XRPUSDC", &wg)
		go BinanceQoute(tokens, "DOGEUSDC", &wg)

		go BinanceQoute(tokens, "ADAUSDC", &wg)
		go BinanceQoute(tokens, "LINKUSDC", &wg)
		go BinanceQoute(tokens, "BCHUSDC", &wg)
		go BinanceQoute(tokens, "LTCUSDC", &wg)

		go BinanceQoute(tokens, "XLMUSDC", &wg)
		go BinanceQoute(tokens, "XMRUSDC", &wg)
		go BinanceQoute(tokens, "XRPUSDC", &wg)
		time.Sleep(1 * time.Second)
	}
	go func() {
		wg.Wait()
		close(tokens)
	}()

	for token := range tokens {
		records = append(records, []string{token.Symbol, token.Price, token.Time, token.Provider})
	}
	fmt.Println("Records:", records)

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(records); err != nil {
		log.Fatal(err)
	}
}

func TradBrainBot() {
	fmt.Println("Welcome to the TradBrain Bot!")

	for range 10 {
		GenerateCSV()
		time.Sleep(1 * time.Second)
	}
}
