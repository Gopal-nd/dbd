package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func loadUser(id int) error {
	if id != 1 {
		return fmt.Errorf("user not found: %w", ErrNotFound)
	}
	return nil
}

type SourceError struct {
	Source string
	Err    error
}

func (e *SourceError) Error() string {
	return fmt.Sprintf("%s: +  %v", e.Source, e.Err)
}

func (e *SourceError) Unwrap() error {
	return e.Err
}

func Err() {
	err := loadUser(2)

	fmt.Println(err)

	fmt.Println(errors.Is(err, ErrNotFound))

	var se *SourceError

	if errors.As(err, &se) {
		fmt.Println(se.Source)
		fmt.Println(se.Err)
		fmt.Println(se.Error())
		fmt.Println(se.Unwrap())

	}
}
