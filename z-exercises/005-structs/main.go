package main

import (
	"fmt"
)

type person struct {
	first   string
	last    string
	flavors []string
}

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

var p1 = person{
	first: "josh",
	last:  "withers",
	flavors: []string{
		"chocolate",
		"vanilla",
	},
}

var p2 = person{
	first: "tanya",
	last:  "gladstone",
	flavors: []string{
		"strawberry",
		"pistacchio",
	},
}

func main() {
	exercise1() // basic struct, with an embedded slice
	exercise2() // embedding a slice in a struct in a map and printing each value
	exercise3() // embedding a struct in a struct, accessing its elements
	exercise4() // anonymous slice of anonymous structs 
}

func exercise1() {
	fmt.Println("------exercise 1------")
	fmt.Println(p1.first, p1.last)
	rangeMaker(p1.flavors)
	fmt.Println(p2.first, p2.last)
	rangeMaker(p2.flavors)
}

func exercise2() {
	fmt.Println("------exercise 2------")
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Printf("%v\n", k)
		fmt.Printf("\tfirst: %v\n", v.first)
		fmt.Printf("\tlast: %v\n", v.last)
		rangeMaker(v.flavors)
	}
}

// helper method to print out each value of a slice of strings
func rangeMaker(flavors []string) {
	for _, v := range flavors {
		fmt.Printf("\t%v\n", v)
	}
}

func exercise3() {
	fmt.Println("------exercise 3------")
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "brickred",
		},
		fourWheel: false,
	}
	
	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "gold",
		},
		luxury: true,
	}
	
	fmt.Println(t)
	fmt.Printf("number of doors: %v\n", t.doors)
	fmt.Println(s)
	fmt.Printf("luxury? : %v\n", s.luxury)
}

func exercise4(){
	fmt.Println("------exercise 4------")
	
	a := []struct{
		name string
		glamorous bool 
		humorous bool 
	}{
		{"josh withers", false, false},
		{"tanya gladstone", true, false},
		{"rupaul queen", true, true},
	}

	for _, v := range a {
		fmt.Printf("%v:\n", v.name)
		fmt.Printf("\tglamorous: %v\n", v.glamorous)
		fmt.Printf("\thumorous: %v\n", v.humorous)
	}
}
