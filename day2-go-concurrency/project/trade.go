package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

func LoadFile(dir string, n int) ([]string, error) {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		if err := GenerateData(dir, n); err != nil {
			return nil, err
		}
	} else if err != nil {
		return nil, err
	}

	res, er := filepath.Glob(filepath.Join(dir, "*.csv"))
	return res, er
}

func GenerateData(dir string, n int) error {
	err := os.MkdirAll(dir, 0o755)
	if err != nil {
		return err
	}

	symbols := []string{"BTCINR", "ETHINR", "SOLINR"}
	basePrice := map[string]float64{"BTCINR": 5800000, "ETHINR": 220000, "SOLINR": 15000}
	sides := []string{"BUY", "SELL"}

	for i := range n {
		var sb strings.Builder
		lines := 1000000 + rand.Intn(151)
		start := time.Now().Add(-24 * time.Hour)

		for j := range lines {
			sym := symbols[rand.Intn(len(symbols))]
			ts := start.Add(time.Duration(j) * time.Minute).Format("2006-01-02T15:04:05")
			qty := rand.Float64() * 2
			price := basePrice[sym] * (0.98 + rand.Float64()*0.04)
			fmt.Fprintf(&sb, "%s,%s,%s,%.4f,%.2f\n", ts, sym, sides[rand.Intn(2)], qty, price)
		}

		name := filepath.Join(dir, fmt.Sprintf("trade_%d.csv", i))
		err := os.WriteFile(name, []byte(sb.String()), 0o644)
		if err != nil {
			return err
		}
	}

	return nil
}

func AnalyzeFile(path string) (map[string]float64, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	volume := make(map[string]float64)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		if len(parts) != 5 {
			continue // skip bad lines
		}
		qty, err := strconv.ParseFloat(parts[3], 64)
		if err != nil {
			continue
		}
		volume[parts[1]] += qty
	}
	return volume, scanner.Err()
}

func AnalyzeFileSlow(path string) (map[string]float64, error) {
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond) // delay
	return AnalyzeFile(path)
}

func PrintTotals(totals map[string]float64) {
	symbols := make([]string, 0, len(totals))
	for s := range totals {
		symbols = append(symbols, s)
	}
	sort.Strings(symbols)

	fmt.Println("\n=== Total volume per symbol ===")
	for _, s := range symbols {
		fmt.Printf("%-8s %10.4f\n", s, totals[s])
	}
}
