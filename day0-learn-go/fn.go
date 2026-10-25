package main

import (
	"errors"
	"fmt"
)

func calculator(a, b int, op string) (float64, error) {
	fa := float64(a)
	fb := float64(b)

	if op == "+" {
		return fa + fb, nil
	} else if op == "*" {
		return fa * fb, nil
	} else if op == "-" {
		return fa - fb, nil
	} else if op == "/" {
		if fb == 0 {
			return 0, fmt.Errorf("divider cannot ba 0")
		} else {
			return fa / fb, nil
		}
	} else {
		return 0, errors.New("not valid")
	}
}
