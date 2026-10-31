package main

import (
	"fmt"
)

func fun() {
	fmt.Println("Hello, World!")

	// array

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 10, 10, 10}

	// find contain duplicate

	val := 9

	fmt.Println(containsDuplicate(arr, val))

	// hashmap

	w1 := "listen"
	w2 := "silent"

	fmt.Println(isAnagram(w1, w2))
}

func containsDuplicate(arr []int, val int) (bool, int) {
	for i, v := range arr {
		if v == val {
			return true, i
		}
	}
	return false, -1
}

func isAnagram(w1 string, w2 string) bool {
	hashmap := make(map[string]int)

	if len(w1) != len(w2) {
		fmt.Println("Not anagrams")
		return false
	}
	for i, v := range w1 {
		hashmap[string(v)] = i
	}

	for _, v := range w2 {
		_, ok := hashmap[string(v)]
		if !ok {
			fmt.Println("Not anagrams")
			return false
		}
	}
	return true
}
