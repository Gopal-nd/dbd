package main

import (
	"errors"
	"fmt"
	"learn_go/fn"
)

func main() {
	fmt.Println("Lets Go!")
	fmt.Print("Is there a new line after the first line?")
	fmt.Printf("I am a %s", "developer")

	var num int = 10
	age := 10
	fmt.Printf("number %d", num)

	str := " Hello i am a string"
	fmt.Printf("name is %s", str)

	sslcPercentage := 99.99
	fmt.Printf("float value %f", sslcPercentage)
	fmt.Printf("float with 1 decimal %.1f", sslcPercentage)

	// binary
	var byteA byte = 'A'
	fmt.Printf(" binary value of 2 is %b %b %b %b %b", 0, 2, int('A'), int('a'), 1) // A = 65
	fmt.Printf("The value of byteA is %b", byteA)
	fmt.Printf("hex of white is %x", "fa1330")
	fmt.Printf("type of ByteA %T", byteA)

	// A =65 ==> 2^6  + 1 = 64 + 1 = 65
	// 100001  (2^)

	canIVote := true
	fmt.Printf(" true is %t", true)

	const pi = 3.14

	const planet = "earth"

	fmt.Print("     ", num, age, str, sslcPercentage, canIVote, pi, planet)

	// functions
	ans, err := fn.Calculator(1, 0, "+")
	// ans, err := calculator(10, 10, "+")

	fmt.Println("")
	if err != nil {
		fmt.Print("There is a error", err)
	}
	fmt.Println(ans)

	// pointers
	myName := "dbd"

	fmt.Println("my name is %p", myName)

	newName := &myName
	fmt.Print(*newName)

	*newName = "distributed backend developer"

	fmt.Println("my newName is %p", *newName)
	fmt.Println("my name is %p", myName)

	value1 := "I am safe"

	println("")

	val := fn.ChangeMe(&value1)
	println(*val)

	println("")
	print(value1)
	println("")

	value2 := "safe"

	fmt.Print(fn.CopyMe(value1))
	println("")
	fmt.Print(value2)

	// cli tax calculator
	v := 0
	tax := 0
	item := 0

	println("** Welcome to the CLI Calculator ***")

	for {
		fmt.Printf("Enter %d item press( 0 or -num ):", item+1)
		temp := 0
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Print("error is %v", err)
		}
		if temp <= 0 {
			break
		}
		v += temp

	}
	for {
		fmt.Printf("Enter Tax percentage):")
		temp := 0
		_, err := fmt.Scan(&temp)
		if err != nil {
			fmt.Print("error is %v", err)
		}
		if temp <= 0 {
			errors.New("value shoud be grater that 0")
			continue
		}
		tax += temp
		break

	}
	fmt.Println(" Your total bill is ", v, "and tax is ", tax)
	fmt.Println("tax is : ", (v / 100 * tax))
	fmt.Print("final payable bill is ", v-(v/100*tax))
}
