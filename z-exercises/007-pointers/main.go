package main

import (
	"fmt"
)

func main() {
	exercise1() // basic pointer
	exercise2() // manipulating values with pointers / addresses
}

func exercise1() {
	fmt.Println("------exercise 1------")
	a := "josh withers"
	fmt.Println("memory address of josh withers: ", &a)
}

func exercise2() {
	fmt.Println("------exercise 2------")
	p := person{"josh", "withers", 14}
	fmt.Println(p)
	changeMe(&p) // pass in the memory address
	fmt.Println(p)
}

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) { // takes in a pointer (memory address)
	*p = person{"tanya", "gladstone", 46} // change the value in memory
	// note the parenthesis around the pointer
	(*p).first = "gatsby" // change again, a specific field in the struct
}
