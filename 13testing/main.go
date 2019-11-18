package main

import "fmt"

func main() {
	fmt.Printf("2 + 3 = %v\n", Sum(2, 3))
	fmt.Printf("137 mod 8 is: %v\n", Modulo(137, 8))
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
