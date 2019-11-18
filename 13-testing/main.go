package main

import (
	"fmt"
	"strings"

	"github.com/SwanHub/Udemy/Golang-ToddMcleod/13-testing/bench2"
	"github.com/SwanHub/Udemy/Golang-ToddMcleod/13-testing/benchmarking"
	"github.com/SwanHub/Udemy/Golang-ToddMcleod/13-testing/examples"
)

func main() {
	fmt.Printf("2 + 3 = %v\n", Sum(2, 3))
	fmt.Printf("137 mod 8 is: %v\n", Modulo(137, 8))
	fmt.Println(examples.Sum2(3, 4, 5))
	benchingCode()           // implementing package benchmarking w function Greet
	benchingCodeComparison() // comparing two work intensive tasks' efficiency
}

// Sum function adds any number of numbers
func Sum(xi ...int) int {
	sum := 0
	for _, v := range xi {
		sum += v
	}
	return sum
}

// Modulo function to find remainder of an integer
func Modulo(x int, y int) int {
	return x % y
}

// Benchmarking code (measuring performance)
func benchingCode() {
	fmt.Println(benchmarking.Greet("josh"))
}

const s = `lkhjsdflgkjnlk sdm;fjallsdc nlksjdhfglks jndflksjndflkjnsgfd
lkhjsdflgkjnlksdm;fjallsdcnlksjdhfglks jndflksjndflkjnsgfd
lkhjsdflgkjnlksdm ;fjallsdcnlksjdhfglksjndflksjndflkjnsgfd
lkhjsdflgkjnlksdm;fjallsdc nlksjdhfglksjndflksjndflkjnsgfd
lkhjsdflgkjnlk sdm;fjallsdcnlksjdhfglksjndflksjndflkjnsgfd
lkhjsdflgkjn lksd m;fjallsdcnlksjdhfglk sjndfl ksjn dflkjn sgfd`

func benchingCodeComparison() {
	xs := strings.Split(s, " ")
	for _, v := range xs {
		fmt.Println(v)
	}

	fmt.Printf("\n%s\n", bench2.Cat(xs))
	fmt.Printf("\n%s\n\n", bench2.Join(xs))
}
