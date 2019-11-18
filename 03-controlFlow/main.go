package main

// Note: we only use the "for" loop keyword in Go. No while or until.

import (
	"fmt"
)

func main() {
	basicFor()
	forAsWhile()
	forWithBreak()
	usingContinue()
	asciiChallenge()
	ifStatement()
	elseStatement()
	moduloOperator()
	switchStatement()
	trueFalseOperators()
}

func basicFor() {
	fmt.Println("-----basic for loop syntax-----")
	// init; condition; post
	for i := 0; i <= 4; i++ {
		fmt.Printf("The outer loop is: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t and the inner loop is: %d\n", j)
		}
	}
}

func forAsWhile() {
	fmt.Println("-----using for like a while loop------")
	x := 0
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("we done it.")
}

func forWithBreak() {
	fmt.Println("------using break in for loop------")
	x := 0
	for {
		if x > 9 {
			fmt.Println("done")
			break
		}
		fmt.Println(x)
		x++
	}
}

func usingContinue() {
	fmt.Println("------using continue keyword-------")
	x := 0
	for {
		x++
		if x > 100 {
			break
		}

		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}

func asciiChallenge() {
	fmt.Println("------using ascii converter------")
	for x := 33; x <= 122; x++ {
		// note the %c converts from decimal/integer to ascii glyph/character
		fmt.Printf("%c\n", x)
	}
}

// new: go includes an initialization statement as an option for if statements
// compiler inputs semi-colons... you could do multiple statements on one line with a slew of semi-colons.
// note: that 'x' is contained to the statement itself. Narrow scope. You couldn't run line 84.
func ifStatement() {
	fmt.Println("------if statement------")
	if x := 42; x == 42 {
		fmt.Println(x)
	}
	// fmt.Println(x)
}

func elseStatement() {
	fmt.Println("------if else statement------")
	x := "john"
	if x == "jake" {
		fmt.Println("it's jake")
	} else if x == "john" {
		fmt.Println("it's john")
	} else {
		fmt.Println("it's nobody")
	}
}

func moduloOperator() {
	fmt.Println("------modulo, for and if statement------")
	for i := 0; i < 20; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func switchStatement() {
	fmt.Println("------using switch statement------")
	switch "Bond" {
	case "Gary", "Cal":
		fmt.Println("false")
	case "Trudy", "Terry", "Bond": // commas are "or" statements
		fmt.Println("true")
		fallthrough // "if this case is true, run the next case as well, regardless "
	case "Melanie":
		fmt.Println("also true")
	default:
		fmt.Println("default case")
	}
}

func trueFalseOperators() {
	fmt.Println("------true false operators------")
	fmt.Printf("true && true\t %v\n", true && 2 == 2)
	fmt.Printf("true && false\t %v\n", true && false)
	fmt.Printf("true || true\t %v\n", true || 2 == 2)
	fmt.Printf("false || true\t %v\n", false || true)
	fmt.Printf("!true\t\t %v\n", !true)
	fmt.Printf("!false\t\t %v\n", !false)
}
