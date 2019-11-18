package main

import (
	"fmt"
)

func main() {
	basicArray()
	basicSlice()
	basicRange()
	slicingSlice()
	appendToSlice()
	deleteSlice()
	usingMake()
	matrices()
	mapping()
	addtoMap()
	deleteValueFromMap()
}

// Note slices are better. Arrays are for memory allocation
func basicArray() {
	fmt.Println("------basic array------")
	// int is referred to as the "element type"
	// the length "5" is part of the TYPE of the array.
	var x [5]int
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)
	l := len(x)
	fmt.Println(l)
}

func basicSlice() {
	fmt.Println("------basic slices------")
	// x := type{values}  composite literal
	x := []int{4, 3, 2, 1, 43}
	fmt.Println(x)
}

func basicRange() {
	fmt.Println("------basic range------")
	x := []int{3, 4, 5, 6, 7}
	for index, value := range x {
		fmt.Println(index, value)
	}
}

func slicingSlice() {
	fmt.Println("------slicing a slice------")
	x := []int{3, 4, 5, 6, 7}
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1:])  // prints out all values starting at position 1
	fmt.Println(x[:3])  // up to but not including the value at position 3
	fmt.Println(x[1:3]) // from position 1, up to but not including position 3

	// alternative method to range
	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

func appendToSlice() {
	fmt.Println("------append values to a slice------")
	x := []int{3, 4, 5, 6, 7}
	fmt.Println(x)
	// append gives back another slice of the same type
	x = append(x, 42, 33, 333)
	fmt.Println(x)

	y := []int{1, 90, 76}
	x = append(x, y...) // "..." unfurls the slice.
	// as opposed to "..." before "y", which means, "take in any number of values"
	// one is "take in any number" || the other is "unfurl out all of the values comma separated"
	fmt.Println(x)
	fmt.Println(y)
}

func deleteSlice() {
	fmt.Println("------delete values from slice (using append)------")
	x := []int{3, 4, 5, 6, 7, 8, 9}
	fmt.Println(x)
	// removes values 5 and 6 in the positions of 2 and 3...
	x = append(x[:2], x[4:]...)
	fmt.Println(x) // destructive

	// dynamic slice assignment
	y := 1
	x = append(x[:y], x[2:]...)
	fmt.Println(x)
}

func usingMake() {
	fmt.Println("------using the make function------")
	x := make([]int, 4, 5)
	// type, length, capacity
	// if you know you have a slice of a certain size, use make for compiler efficiency
	// length is the current slice size... capacity determines the underlying array size
	// 4 and 5 are like the initial range conditions for the slice you're working with
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) // capacity

	x[0] = 5
	x[3] = 12
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	// when you exceed the capacity, the underlying array is deleted, copied over and doubled in size
	x = append(x, 42, 53, 64)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}

func matrices() {
	fmt.Println("------multi-dimensional slices (matrices)------")
	row1 := []string{"this", "is", "first"}
	row2 := []string{"and", "dis", "second"}
	matrix := [][]string{row1, row2}
	fmt.Println(matrix)
}

func mapping() {
	fmt.Println("------using maps (k/v store)------")
	// SUPER FAST / EFFICIENT lookup
	m := map[string]int{
		"james": 43,
		"cory":  22, // need trailing comma
	}
	fmt.Println(m)
	fmt.Println(m["james"])
	fmt.Println(m["unknown"]) // zero value return for no key declared
	// checking for existence "comma ok"
	v, ok := m["unknown"]
	fmt.Println(v)
	fmt.Println(ok)

	// integrating to conditional
	if v, ok = m["james"]; ok {
		fmt.Printf("Successful value: %v\n", v)
	}
}

func addtoMap() {
	fmt.Println("------adding values to maps------")
	m := map[string]int{
		"james": 43,
		"cory":  22, // need trailing comma
	}
	// adding values
	m["new"] = 45
	fmt.Println(m)

	// ranging over key values
	for key, value := range m {
		fmt.Println(key, value)
	}
}

func deleteValueFromMap() {
	fmt.Println("------adding values to maps------")
	m := map[string]int{
		"cal":     42,
		"miranda": 32,
	}

	fmt.Println(m)

	delete(m, "cal")
	fmt.Println(m)
}
