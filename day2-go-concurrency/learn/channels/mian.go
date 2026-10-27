package channels

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"sync"
)

type Result struct {
	Symbol string `json:"symbol"`
	Price  string `json:"price"`
}

func GetTheSOlPrice(ch chan string) {
	resp, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol=SOLUSDC")
	if err != nil {
		fmt.Print(err)
		ch <- "err"

	}

	data, _ := io.ReadAll(resp.Body)
	var result Result
	if err := json.Unmarshal(data, &result); err != nil {
		fmt.Print(err)
		ch <- err.Error()
		return
	}
	ch <- result.Price + strconv.Itoa(rand.Intn(100))
	// fmt.Println("function ended")
}

func Producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := range 2 {
		ch <- i
	}
}

func ProducerClose(ch chan int) {
	for i := range 2 {
		ch <- i
	}

	// sender should close the channel, never the receiver (if reciver try it will panic)
	close(ch)
}

func Chan() {
	fmt.Println("Go Channels")

	ch := make(chan string)

	go GetTheSOlPrice(ch)
	go GetTheSOlPrice(ch)
	go GetTheSOlPrice(ch)
	go GetTheSOlPrice(ch)
	go GetTheSOlPrice(ch)
	go GetTheSOlPrice(ch)
	go GetTheSOlPrice(ch)

	// for v := range ch {
	// 	fmt.Println(v)
	// }

	for range 7 {
		fmt.Print(<-ch)
	}
	go func() {
		ch <- "test"
	}()
	fmt.Println(<-ch)

	ch2 := make(chan int, 40)
	var wg sync.WaitGroup
	wg.Add(3)
	go Producer(ch2, &wg)
	go Producer(ch2, &wg)
	go Producer(ch2, &wg)
	func() {
		wg.Wait()
		ch2 <- 999999
		close(ch2)
	}()
	for v := range ch2 {
		fmt.Println(v)
	}

	ch3 := make(chan int)
	go ProducerClose(ch3)
	for v := range ch3 {
		fmt.Println(v)
	}

	// ch4 := make(chan int)

	// ch4 <- 100

	// err, ok := <-ch4
	// if ok {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println("error: ", err)
	// }
}
