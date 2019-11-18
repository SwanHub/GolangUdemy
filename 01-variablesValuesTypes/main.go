package main

import "fmt"

var x = 42

// illegal ... x := 42
var s string

type legend string

func main() {
	fmt.Println("Whatup")
	foo()

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("even number: ", i)
		}
	}

	n, err := fmt.Println("Catch my number of bytes and the potential error")
	fmt.Println(n)
	fmt.Println(err)

	_, newErr := fmt.Println("don't print the bytes, still print error if occurs")
	fmt.Println(newErr)

	fmt.Println(x)
	// Printf == print something with a specific format (interpolation)
	// Sprintf / Sprintln == Print to a string. Not just standard output.
	fmt.Printf("%T\n", x) // prints the type///
	s := fmt.Sprintf("%v\n", x)
	fmt.Println(s)

	var a legend = "legend"
	fmt.Println(a)

	s = string(a)
	fmt.Printf("\t%s\n", s)
}

func foo() {
	fmt.Println("bar")
}
