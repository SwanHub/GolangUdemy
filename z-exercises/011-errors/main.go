package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	exercise1() // basic error handle
	exercise2() // using fmt.Errorf
	exercise3() // creating a customErr struct which implements the builtin.error interface
	exercise4() // using the error type in a custom error struct
}

type person struct {
	First   string
	Last    string
	Sayings []string
}

func exercise1() {
	fmt.Println("------exercise 1------")
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := json.Marshal(p1)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(bs))
}

func exercise2() {
	fmt.Println("------exercise 2------")
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}

	bs, err := toJSON(p1)
	if err != nil {
		log.Println("you done messed it all up, error: ", err)
		return
	}

	fmt.Println(string(bs))

}

func toJSON(a interface{}) ([]byte, error) {
	bs, err := json.Marshal(a)
	if err != nil {
		return []byte{}, fmt.Errorf("no bueno, amigo - json bad: %v", err)
		// return []byte{}, errors.New(fmt.Sprintf("no bueno, amigo - json bad: %v", err))
	}
	return bs, nil
}

func exercise3() {
	fmt.Println("------exercise 3------")
	// third to print
	ce := customErr{"this ain't good"}
	printMyErr(ce)
}

type customErr struct {
	info string
}

func (c customErr) Error() string {
	// second to print
	return fmt.Sprintf("your error with a side of sauce, %v", c.info)
}

// WE CAN PASS THIS IN AS ERROR BC OF TYPE INTERFACE going on behind the scenes
// Because ce has a method Error(), we can pass it into the error function and get buckets therein
func printMyErr(e error) { // 1st to print
	fmt.Printf("bad shit, %v\n", e) // conversion
	// fmt.Printf("bad shit, %v\n", e.(customErr).info) // assertion
}

type sqrtError struct {
	lat  string
	long string
	err  error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v %v", se.lat, se.long, se.err)
}

func exercise4() {
	fmt.Println("------exercise 4------")
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		se := fmt.Errorf("this is an error, you can't use this number: %v", f)
		return 0, sqrtError{"FAKE LAT", "FAKE LONG", se}
	}
	return 42, nil
}

// Add one number to itself aka multiply by two
func Add(i int) int {
	return i + i
}
