package main

import (
	"fmt"
)

func main() {
	fmt.Println("pointers")
	basicPointer() // basic pointer declaration
	morePointer()
}

// memory is a post office with PO Boxes
// all values have a memory addres
// advantage of using a pointer is you can just pass the ADDRESS as opposed to VALUE

func basicPointer() {
	fmt.Println("------basic pointers------")
	a := 4233
	fmt.Println(a)
	fmt.Println(&a) // displays the address in memory

	fmt.Printf("%T\n", a)  //
	fmt.Printf("%T\n", &a) // *int (type) &a (value) *a('*'is operator --> gives value stored at address)

	var b *int = &a
	fmt.Println(b)  // pointer to memory address
	fmt.Println(*b) // dereference, gives value stored at address when b is an address
	fmt.Println(*&b)

	*b = 43331
	fmt.Println(a) // b and a are intertwined in this situation, bc same memory address

}

func morePointer() {
	fmt.Println("------more pointers------")
	x := 0
	fmt.Println("x before: ", &x)
	fmt.Println("x before: ", x)
	helper(&x)
	fmt.Println("x after: ", &x)
	fmt.Println("x after: ", x)
}

func helper(y *int) {
	fmt.Println("y before: ", y)
	fmt.Println("y before: ", *y)
	*y = 44
	fmt.Println("y after: ", y)
	fmt.Println("y after: ", *y)
}

// methods sets (all methods attached to a type)
