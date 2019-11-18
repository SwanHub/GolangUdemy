package main

import (
	"fmt"
)

func main() {
	exercise1()  // basic for loop
	exercise2()  // nested loop with type conversion
	exercise3()  // for loop with one condition
	exercise4()  // for loop with no conditions.
	exercise5()  // modulo operator
	exercise6()  // if statement
	exercise7()  // if, else if, else statement
	exercise8()  // switch statement no switch expression
	exercise9()  // switch statement with switch expression
	exercise10() // true false statements
}

func exercise1() {
	fmt.Println("--------exercise 1---------")
	for i := 0; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func exercise2() {
	fmt.Println("")
	fmt.Println("--------exercise 2---------")
	for i := 65; i <= 90; i++ {
		fmt.Printf("%v\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}

func exercise3() {
	fmt.Println("")
	fmt.Println("--------exercise 3---------")
	a := 1995
	for a < 2020 {
		fmt.Println(a)
		a++
	}
	fmt.Printf("Total of %v years\n", a-1995)
}

func exercise4() {
	fmt.Println("")
	fmt.Println("--------exercise 4---------")
	a := 1995
	for {
		if a > 2019 {
			break
		}
		fmt.Println(a)
		a++
	}
}

func exercise5() {
	fmt.Println("")
	fmt.Println("--------exercise 5---------")
	for i := 10; i <= 100; i++ {
		fmt.Printf("%v divide by 4, remainder %v\n", i, i%4)
	}
}

func exercise6() {
	fmt.Println("")
	fmt.Println("--------exercise 6---------")
	x := 44
	if x == 44 {
		fmt.Println("my if statement is working")
	}
}

func exercise7() {
	fmt.Println("")
	fmt.Println("--------exercise 7---------")
	x := 44
	if x == 42 {
		fmt.Println("we are 42")
	} else if x == 41 {
		fmt.Println("we are actually 41")
	} else {
		fmt.Printf("the right answer was %v\n", x)
	}
}

func exercise8() {
	fmt.Println("")
	fmt.Println("--------exercise 8---------")
	switch {
	case false:
		fmt.Println(false)
	case true:
		fmt.Println(true)
	}
}

func exercise9() {
	fmt.Println("")
	fmt.Println("--------exercise 9---------")
	var favSport string = "basketball"
	switch favSport {
	case "baseball", "football", "polo":
		fmt.Println("all wrong")
	case "basketball":
		fmt.Println("swoosh")
	}
}

func exercise10() {
	fmt.Println("")
	fmt.Println("--------exercise 10---------")
	fmt.Printf("true && true evaluates to \t%v\n", 2 == 2 && true)
	fmt.Printf("true && false evaluates to \t%v\n", true && false)
	fmt.Printf("true || true evaluates to \t%v\n", 2 == 2 || true)
	fmt.Printf("false || true evaluates to \t%v\n", false || true)
	fmt.Printf("!true evaluates to \t\t%v\n", !true)
	fmt.Printf("!false evaluates to \t\t%v\n", !false)
}
