package main

import (
	"fmt"
	"sync"
	"time"
)

const numWorkers = 4

type Result struct {
	File   string
	Volume map[string]float64
	Err    error
}

func worker(id int, jobs <-chan string, book *VolumeBook, wg *sync.WaitGroup) {
	defer wg.Done()
	for path := range jobs { // keeps going until jobs is closed
		vol, err := AnalyzeFile(path)
		if err != nil {
			fmt.Println(path, "error:", err)
			continue
		}
		book.Add(vol)
	}
}

type VolumeBook struct {
	mu     sync.Mutex
	totals map[string]float64
}

func (b *VolumeBook) Add(vol map[string]float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for sym, qty := range vol {
		b.totals[sym] += qty
	}
}

func main() {
	start := time.Now()
	fmt.Println("Welcome to the project!")
	files, err := LoadFile("data", 10)
	if err != nil {
		fmt.Println("Error loading files:", err)
		return
	}
	// fmt.Println("Files:", files)

	var wg sync.WaitGroup
	jobs := make(chan string, len(files))
	book := &VolumeBook{totals: make(map[string]float64)}

	// fan out
	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, book, &wg)
	}

	go func() {
		for _, f := range files {
			jobs <- f
		}
		close(jobs)
	}()

	// fan in

	wg.Wait()

	fmt.Println("Total volume:", book.totals)

	defer fmt.Println("Time taken:", time.Since(start))
}
