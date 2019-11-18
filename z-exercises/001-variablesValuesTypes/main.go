package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func exercise1() {
	fmt.Println("--------exercise 1 --------")
	fmt.Printf("\t%v\n", x)
	fmt.Printf("\t%v\n", y)
	fmt.Printf("\t%v\n", z)
	fmt.Println(x, y, z)
}

var j int
var k string
var l bool

func exercise2() {
	fmt.Println("--------exercise 2---------")
	fmt.Printf("\t%T\n", j)
	fmt.Printf("\t%T\n", k)
	fmt.Printf("\t%T\n", l)

	fmt.Println(j)
	fmt.Println(k)
	fmt.Println(l)
}

func exercise3() {
	fmt.Println("--------exercise 3---------")
	s := fmt.Sprintf("%d // %s // %t", x, y, z)
	fmt.Println(s)
}

func exercise45() {
	type ching int
	var x ching
	fmt.Println("--------exercise 4---------")
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	fmt.Println("-exercise 5--/conversion/--")
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}

func main() {
	fmt.Println("Whaup")

	exercise1()
	exercise2()
	exercise3()
	exercise45()

}
