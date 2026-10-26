package main

import (
	"fmt"
	"go_continue/crypto"
	"maps"
	"os"
	"slices"
	"sync"
	"time"
)

func main() {
	//arrays

	arr := [10]int{1, 2, 3, 4, 5, 6}
	fmt.Print("%v", arr)

	arr1 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("")
	fmt.Printf("array is : %+v", arr1)

	//array opperations

	arr1[0] = 10000
	fmt.Println("\n")
	fmt.Printf("array is : %+v", arr1)

	//push at lsat

	arr1 = append(arr1, 1000)
	fmt.Println("\n")
	fmt.Printf("array is : %+v", arr1)

	// pop at last
	arr1 = arr[:len(arr1)-1]
	fmt.Print("\n")
	fmt.Print(arr1)

	fmt.Print(arr1)
	//loop
	fmt.Print("\n")
	for Index, value := range arr1 {
		fmt.Print(Index, "->", value, "\n")
	}

	//append at front
	arr1 = append([]int{100000}, arr1...)
	fmt.Println("\n")
	fmt.Printf("array is : %+v", arr1)

	//remove the front one
	arr1 = arr1[1:]
	fmt.Println("\n")
	fmt.Println(arr1)

	arr2 := make([]int, len(arr1))

	copy(arr2, arr1)
	fmt.Printf("coping value of array :%v", arr2)
	fmt.Println(cap(arr1))
	fmt.Println(cap(arr))

	fmt.Println(len(arr1))
	fmt.Println(len(arr))

	// MAP data type
	var map1 = map[int]string{
		1: "goapl",
		2: "verdhan",
	}

	var map2 = make(map[string]int)

	map2["Gopal"] = 1
	map2["Shashank"] = 2

	fmt.Printf("%+v", map1)
	fmt.Print(map1)
	fmt.Print(map2)

	for key, value := range map2 {
		fmt.Println(key, "->", value)
	}
	delete(map2, "Shashank")
	for key, value := range map2 {
		fmt.Println(key, "->", value)
	}
	// check if exist

	val, exist := map2["Gopal"]
	if exist {
		fmt.Printf("%d", val)
	}
	fmt.Println(exist)

	var nilMap map[string]int
	fmt.Println(nilMap == nil) // true
	println(nilMap)
	fmt.Println(nilMap["key"]) // 0 (Safe to read)

	// nilMap["key"] = 10// error

	correctMap := make(map[int]Student)

	correctMap[1] = Student{
		name:   "gopal",
		age:    21,
		gender: str("male"),
	}

	correctMap[2] = Student{
		name:   "rgv",
		age:    12,
		gender: str("male"),
	}

	for i, v := range correctMap {
		fmt.Println(i, v.age, *v.gender)
	}

	// no ternary operator or the distructuring in the go

	// switch

	day := 41
	switch day {
	case 0:
		{
			print("sunday")
		}
	case 1:
		{
			print("monday")
		}

	case 2:
		{
			print("tuesday")
		}
	case 3:
		{
			print("wednesday")
		}
	case 4:
		{
			print("thursday")
		}
	case 5:
		{
			print("friday")
		}
	case 6:
		{
			print("saturday")
		}
	default:
		{
			print("wrong number")
		}

	}

	// defer retrun before the function return
	// 1. close db, file
	go func() {
		data, err := os.ReadFile("go.mod")
		da, _ := os.Open("go.mod")

		fmt.Print(" I am working **********************")
		defer da.Close()
		if err != nil {
			println(err)
			return
		}
		fmt.Println(string(data))
	}()
	time.Sleep(500 * time.Millisecond)
	// 2. calculate the time difference of the start and end of the process
	func() {
		// _ := time.Now()
		temp := 0
		for val := range 1_000_000_00 {
			temp += val
		}

		defer fmt.Printf("value is : %T %d", temp, temp)

	}()

	fun := [80_000_000]int{1, 3, 4, 5}
	fmt.Print(fun[:4])

	// time.Sleep(10 * time.Second)
	// 3. helps in locking and unloacking
	store := &Store{
		num: 0,
	}

	fmt.Println(store.num)
	store.Increment()
	fmt.Println(store.num)
	store.Increment()
	fmt.Println(store.num)
	store.Increment()
	fmt.Println(store.num)
	// 4. recover from the panic
	go func() {
		defer func() {
			rec := recover()
			if rec != nil {
				fmt.Printf("from the panic")
			}
		}()

		panic("some  thing is wrong")
	}()

	time.Sleep(1 * time.Second)
	fmt.Print("Execution is done")

	s1 := Student{
		name:   "rahul",
		age:    10,
		gender: nil,
	}
	fmt.Println(s1.name)
	ChangeME(&s1)
	fmt.Println(s1.name)

	s2 := Student{
		name:   "go",
		age:    17,
		gender: nil,
	}
	fmt.Println(s2.age)
	CopyOfME(s2)
	s2.IncrementMyAge()
	fmt.Println(s2.age)

	arr0 := []int{0, 1, 2, 3, 4, 5, 6}

	fmt.Print(arr0)
	arr0 = append(arr0[:3], arr0[4:]...)

	fmt.Print(arr0)

	// project.Taskers()

	numbers := []int{10, 15, 12, 5}

	fmt.Println(Sum(numbers))

	//callback functions
	removeEven := func(n int) bool {
		if n%2 == 0 {
			return false
		} else {
			return true
		}
	}

	fmt.Println(Filter(numbers, removeEven))
	slices.Reverse(numbers)
	remove10and5 := func(s int) bool {

		if s == int(10) || s == int(5) {
			return true
		} else {
			return false

		}

	}

	fmt.Println(numbers)
	fmt.Println(slices.DeleteFunc(numbers, remove10and5))
	fmt.Println(numbers)

	for i := range maps.Keys(map1) {
		fmt.Println(i)
	}

	crypto.FindArb()
}

type Number interface{ ~int | ~int64 | ~float64 }

func Sum[T Number](nums []T) T {
	var total T
	for _, v := range nums {
		total += v
	}
	return total
}

// keep is a callback function in go
func Filter[T Number](nums []T, keep func(T) bool) []T {
	var ans []T
	for _, val := range nums {
		if keep(val) {
			ans = append(ans, val)
		}

	}
	return ans
}

func CopyOfME(s Student) {
	s.name = "kumar"
}

func ChangeME(s *Student) {
	s.name = "kumar"
}

func str(s string) *string {
	return &s
}

type Store struct {
	mu  sync.Mutex
	num int
}

func (s *Store) Increment() {
	s.mu.Lock()
	s.num += 1
	defer s.mu.Unlock()
}

type Student struct {
	name   string
	age    int
	gender *string
}

func (s Student) Summary() {
	fmt.Println("name is ", s.name, "age is ", s.age, "gender", s.gender)
}

func (s *Student) IncrementMyAge() {
	s.age += 1
}
