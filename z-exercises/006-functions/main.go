package main

import (
	"fmt"
	"math"
)

func main() {
	exercise1()  // single and multiple returns
	exercise2()  // variadic parameters
	exercise3()  // defer
	exercise4()  // methods
	exercise5()  // interfaces
	exercise6()  // anonymous function
	exercise7()  // anonymous function saved in var
	exercise8()  // return function from another function
	exercise9()  // callback function
	exercise10() // closure
}

func exercise1() {
	fmt.Println("------exercise 1------")
	fmt.Println(foo1())
	fmt.Println(bar1())
}

func foo1() int {
	return 4
}

func bar1() (int, string) {
	return 14, "nice"
}

func exercise2() {
	fmt.Println("------exercise 2------")
	foo2([]int{1, 2, 3, 4}...)
	bar2([]int{1, 2, 3, 4})
}

func foo2(x ...int) {
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
}

func bar2(x []int) {
	sum := 0
	for _, v := range x {
		sum += v
	}
	fmt.Println(sum)
}

func exercise3() {
	fmt.Println("------exercise 3------")
	defer func() { fmt.Println("first") }()
	func() { fmt.Println("second") }()
}

func exercise4() {
	fmt.Println("------exercise 4------")
	p1 := person{
		first: "josh",
		last:  "withers",
		age:   14,
	}
	p1.speak()
}

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println(p.first, p.last, p.age)
}

func exercise5() {
	fmt.Println("------exercise 5------")
	c := circle{
		radius: 3.4,
	}
	info(c)
	sq := square{
		length: 4,
		width:  5,
	}
	info(sq)
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

type square struct {
	length float64
	width  float64
}

func (s square) area() float64 {
	return s.length * s.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func exercise6() {
	fmt.Println("------exercise 6------")
	func() {
		fmt.Println("anonymous function ftw")
	}()
}

func exercise7() {
	fmt.Println("------exercise 7------")
	f := func() {
		fmt.Println("anonymous #2 function ftw")
	}
	f()
}

func exercise8() {
	fmt.Println("------exercise 8------")
	e8helper()()
}

func e8helper() func() {
	return func() {
		fmt.Println("inception functions")
	}
}

func exercise9() {
	fmt.Println("------exercise 9------")

	e9helper(func() { fmt.Println("double dose of inception") })()
}

func e9helper(f func()) func() {
	return f
}

func exercise10() {
	fmt.Println("------exercise 10------")

	f := e10helper()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}

func e10helper() func() int {
	sum := 0
	return func() int {
		sum++
		return sum
	}
}
