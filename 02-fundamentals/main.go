package main

import (
	"fmt"
	"runtime"
)

// No triple equal sign in Go.
func main() {
	// bools
	a := 7
	b := 9

	if a == b {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	// int
	x := 22
	y := 22.345345
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	// runtime
	fmt.Println(runtime.GOOS)   // Go operating system
	fmt.Println(runtime.GOARCH) // amd64 (64 bit processor)

	// string
	s := `"Heyo"`
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bs := []byte(s) // slice of uint8 (reminder: uint8 is #s 0 - 255 or --2^8)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i]) // UTF-8 code point
	}

	for i, v := range s {
		fmt.Printf("at index: %d, we have hex: %#x\n", i, v) // hexadecimal
	}

	bitShifting()

}

/////////// NOTES
// Use int and float64
// byte == uint8 (8 digits == 2^8 == 256 combinations)
// rune == int32
// each rune is a codepoint in utf-8

func numeralWork() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s) // byte array of the string, written in decimal.
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)  // type for 72
	fmt.Printf("%b\n", n)  // binary for 72
	fmt.Printf("%#X\n", n) // hexadecimal for 72 "0X" means "we're using hexademical"
}

// Constants
func constantWork() {

	// shorthand to declare multiple constants.
	const (
		a        = 42
		b        = 42.78
		c        = "Cory Booker" // untyped constant
		d string = "the dude"    // typed constant
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Printf("%T\n", a) // picks up the type automatically
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)

}

// Iota, auto-incrementation type... for constants.
func iotaWork() {

	const (
		a = iota
		b = iota
		c = iota
	)

	// resets the count in a new constant set
	const (
		d = iota
		e = iota
		f = iota
	)

	fmt.Println(a) // 0
	fmt.Println(b) // 1
	fmt.Println(c) // 2
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", a) // they're all of type int

}

// Bit Shifting (iota used for this)
func bitShifting() {
	x := 1
	fmt.Printf("%d\t\t%b\n", x, x) // %d is decimal, %b is binary

	y := x << 3                    // bit shift 3 to the left.
	fmt.Printf("%d\t\t%b\n", y, y) // bit shifting multiplies by two. Same as in decimal... adding a zero.

	const (
		_  = iota // a
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10) // bit shifting by ten...
	)

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb) // shifts ten each time.

}
