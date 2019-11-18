package main

import (
	"fmt"
)

func main() {
	exercise1()
	exercise2()
	exercise3()
	exercise4()
	exercise5()
	exercise6()
}

// printing decimal, binary, hex
func exercise1() {
	fmt.Println("--------exercise 1---------")
	x := 4

	fmt.Printf("%d\t%b\t%#X\n", x, x, x)

}

// conditionals
func exercise2() {
	fmt.Println("")
	fmt.Println("--------exercise 2---------")
	var a, b, c, d, e, f bool

	a = (5 == 4)
	b = (5 <= 4)
	c = (5 >= 4)
	d = (5 != 4)
	e = (5 < 4)
	f = (5 > 4)

	fmt.Println(a, b, c, d, e, f)
}

// constants
func exercise3() {
	fmt.Println("")
	fmt.Println("--------exercise 3---------")

	const (
		a int    = 5
		b        = 5
		c string = "Jerry"
		d        = "John"
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// bit shifting
func exercise4() {
	fmt.Println("")
	fmt.Println("--------exercise 4---------")

	a := 5
	fmt.Printf("%d\t%b\t%#X\n", a, a, a)
	b := a << 1
	fmt.Printf("%d\t%b\t%#X\n", b, b, b)
}

// string literal
func exercise5() {
	fmt.Println("")
	fmt.Println("--------exercise 5---------")
	var s string = `John
	is a cool dude, yeah?`

	fmt.Println(s)
}

func exercise6() {
	fmt.Println("")
	fmt.Println("--------exercise 6---------")
	const (
		a = 2020 + iota
		b = 2020 + iota
		c = 2020 + iota
		d = 2020 + iota
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

// miscellaneous: bit is a "binary digit"
