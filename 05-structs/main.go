package main

import (
	"fmt"
)

type person struct {
	first string
	last  string
}

type secretAgent struct {
	person
	first string
	ltk   bool
}

// aggregate data type, composition of different types.
// creating values of a certain customized type
func main() {
	basicStruct()     // basic struct literal
	embeddedStruct()  // struct within a struct, handling name collision
	anonymousStruct() // anonymous struct + anonymous slice of structs
}

func basicStruct() {
	fmt.Println("-----basic struct-----")
	p1 := person{
		first: "james",
		last:  "bond",
	}

	p2 := person{
		first: "josh",
		last:  "withers",
	}

	fmt.Println(p1, p2)
	fmt.Println(p1.first, p1.last)
	fmt.Println(p2.first, p2.last)
}

func embeddedStruct() {
	fmt.Println("-----embedded struct-----")
	sa1 := secretAgent{
		person: person{
			first: "james",
			last:  "bond",
		},
		first: "name collision",
		ltk:   true,
	}

	// person (inner type) is promoted to the outer type (secretAgent)
	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.person.first, sa1.last, sa1.ltk)
}

func anonymousStruct() {
	fmt.Println("-----anonymous struct-----")
	person := struct {
		first string
		last  string
	}{
		first: "kendall",
		last:  "jenner",
	}

	people := []struct {
		first string
		last  string
	}{
		{"ian", "mckellen"},
		{"george", "lucas"},
	}

	fmt.Println(person)
	fmt.Println(people)
}
