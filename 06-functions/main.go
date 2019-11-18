package main

import (
	"fmt"
)

// function syntax: func (r receiver) identifier(params) (return(S)) { ... }

func main() {
	noParams()
	oneParam("single parameter")
	r := returnValue("return this")
	fmt.Println(r)
	s, b := twoReturns("josh", "withers")
	fmt.Printf("name: %v\n\tlonger than 10 characters? %v\n", s, b)
	variadicParams(2, 3, 4, 5, 6)
	unfurlSlice()
	deferFunction()
	usingMethods()
	// interfaceFunction()
	anonymousFunction()
	functionExpression()
	doubleFunc()
	fmt.Println("------callback function------")
	callbackFunctions(2, 3, sum)
	callbackFunctions(2, 3, multiply)
	codeBlockInCodeBlock()
	demonstratingClosure()
	fmt.Println("------recursive function------")
	fmt.Printf("factorial of 4 is: %v\n", factorial(4))
	fmt.Println("------loop factorial function------")
	fmt.Printf("factorial of 4 is: %v\n", loopFactorial(4))
}

func noParams() {
	fmt.Println("------bare function------")
	fmt.Println("no params here")
}

// everything in Go is PASSED BY VALUE (not reference)
func oneParam(s string) {
	fmt.Println("------one param function------")
	fmt.Println(s)
}

func returnValue(s string) string {
	fmt.Println("------return one value function------")
	r := fmt.Sprintf("%v value", s)
	return r
}

func twoReturns(first, last string) (string, bool) {
	fmt.Println("------return two values function------")
	name := fmt.Sprintf("%v %v", first, last)
	if len(name) > 10 {
		return name, true
	}
	return name, false
}

// slice of int is the type of a variadic parameter
func variadicParams(x ...int) {
	fmt.Println("------variadic params function------")
	fmt.Println(x)
	fmt.Printf("variadic type: %T\n", x)

	sum := 0
	for i, v := range x {
		fmt.Printf("for el in pos %v, add %v to the total %v\n", i, v, sum)
		sum += v
	}
	fmt.Printf("The total is: %v\n", sum)
}

func unfurlSlice() {
	fmt.Println("------unfurl slice function------")

	x := []int{2, 3, 4}
	sum := quickSum(x...)
	fmt.Printf("Sum of slice values is: %v\n", sum)

}

//helper method to "unfurlSlice" function
func quickSum(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func deferFunction() {
	fmt.Println("------defer function------")
	defer first()
	second()
}

func first() {
	fmt.Println("first")
}

func second() {
	fmt.Println("second")
}

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	ltk bool
}

// requires struct -- this is where we use the "r receiver" part of func syntax
func usingMethods() {
	fmt.Println("------using methods-function------")

	sa1 := secretAgent{
		person: person{
			first: "josh",
			last:  "withers",
		},
		ltk: true,
	}

	p1 := person{
		first: "tanya",
		last:  "gladstone",
	}

	fmt.Println(sa1)
	sa1.speak()

	fmt.Println("------using interface design------")
	interfaceFunction(sa1)
	sa1.speak()
	interfaceFunction(p1)
	p1.speak()

}

// attaches the speak function to secretAgents
func (s secretAgent) speak() {
	fmt.Printf("my name is %v %v\n", s.first, s.last)
}

func (p person) speak() {
	fmt.Printf("my name is %v %v\n", p.first, p.last)
}

type human interface {
	speak()
}

// a value can be of more than one type. sa1 will be of type human and secretAgent
func interfaceFunction(h human) {
	switch h.(type) {
	case secretAgent:
		fmt.Printf("I'm a human (secret agent), here are my credentials: %v\n", h)
		fmt.Printf("\tand let me prove it: %v\n", h.(secretAgent).ltk)
	case person:
		fmt.Printf("I'm a human (normal person), here are my credentials: %v\n", h)
		fmt.Printf("\tand let me prove it: %v\n", h.(person).first)
	}
}

func anonymousFunction() {
	fmt.Println("------using anonymous function------")
	func() { // no params
		fmt.Println("absolute beauty, josh withers")
	}()
	func(x int) {
		fmt.Printf("anonymous function with parameter: %v\n", x)
	}(4)
}

// functions are first class citizens
func functionExpression() {
	fmt.Println("------function expression------")
	f := func() {
		fmt.Println("I saved this in a variable, then called it. Crazy?")
	}
	f()
}

func doubleFunc() {
	fmt.Println("------returning functions within functions------")
	fmt.Printf("A good book is Fahrenheit %v\n", returnFuncFromFunc()())
	fmt.Printf("Type of return is: %T\n", returnFuncFromFunc())
}

// a function is a TYPE, so we can return it easy.
func returnFuncFromFunc() func() int {
	return func() int {
		return 451
	}
}

func callbackFunctions(x, y int, f func(x int, y int) int) {
	fmt.Printf("The result of this function is: %v\n", f(x, y))
}

func sum(x int, y int) int {
	return x + y
}

func multiply(x int, y int) int {
	return x * y
}

// not practical, just demonstration
func codeBlockInCodeBlock() {
	fmt.Println("------code block inside code block------")
	y := 23
	fmt.Printf("outer block value: %v\n", y)
	{
		y := 24
		fmt.Printf("\tinner block value: %v\n", y)
	}
}

func demonstratingClosure() {
	fmt.Println("------closure function------")
	a := incrementor()
	b := incrementor()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

// enclosing a variable in some code, so its scope is limited just to that code.
func incrementor() func() int {
	var x int // set to zero value
	return func() int {
		x++
		return x
	}
}

// recursion is unnecessarily complex, not as good as loops
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func loopFactorial(n int) int {
	product := 1
	for ; n > 0; n-- {
		product *= n

	}
	return product
}
