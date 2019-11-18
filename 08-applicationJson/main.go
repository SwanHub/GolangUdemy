package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io"
	"os"
	"sort"
)

func main() {
	marshalingJSON()     // taking struct // slice // map and creating json (creating json string)
	unmarshalingJSON()   // taking []bytes and making it a struct (interpreting json string)
	writerInterface()    // under the hood of interfaces and how they work with Writers
	sortingSliceValues() // Sort
	customSort()         // Sort on your own terms
	bcryptfunction()     // encrypting passwords
}

func marshalingJSON() {
	p1 := person{"josh", "withers", 14}
	p2 := person{"tanya", "gladstone", 46}

	people := []person{p1, p2}
	fmt.Println(people)

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error: ", err)
	}
	// Caution :: Note :: won't work if the struct fields are lower case
	fmt.Println(string(bs))
}

type person struct {
	First string `json:"First"`
	Last  string `json:"Last"`
	Age   int    `json:"Age"`
}

func unmarshalingJSON() {
	fmt.Println("------unmarshaling JSON------")
	// our json
	s := `[{"First":"josh","Last":"withers","Age":14},{"First":"tanya","Last":"gladstone","Age":46}]`
	// note this is how we convert a string to a slice of bytes
	bs := []byte(s)

	fmt.Printf("%T\n", s)  // string
	fmt.Printf("%T\n", bs) // []uint8 (byte is alias of uint8) (:: this is a slice of bytes)

	var people []person

	// takes in the memory address by documentation definition
	// Why necessary to have memory pointer? Probably to change value directly and permanently
	// hence why I can then use people later on with the changed value from this function
	err := json.Unmarshal(bs, &people)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println(people) // value of type struct
	for i, v := range people {
		fmt.Println("PERSON NUMBER ", i)
		fmt.Printf("\tCredentials: %v %v %v\n", v.First, v.Last, v.Age)
	}
}

// Encode and Decode operate directly, skipping middle man of creating a variable like in marshaling
// this is what you use for web development, encoding and decoding
func writerInterface() {
	fmt.Println("------using Writer Interface------")
	fmt.Println("standard out with println")
	fmt.Fprintf(os.Stdout, "standard out with fprintf, which takes a file, in this case os.Stdout\n")
	io.WriteString(os.Stdout, "standard out with io.WriteString, which also takes in a file os.Stdout\n")
}

func sortingSliceValues() {
	fmt.Println("------sorting values from slices------")
	i := []int{12, 17, 14, 18, 5}
	s := []string{"josh", "withers", "tanya", "gladstone"}

	fmt.Println(i)
	sort.Ints(i) // pull in sort package to do this work
	fmt.Println(i)
	// what it takes to reverse a slice of integers
	sort.Sort(sort.Reverse(sort.IntSlice(i)))
	fmt.Println(i)

	fmt.Println(s)
	sort.Strings(s)
	fmt.Println(s)
	// what it takes to reverse a slice of strings
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
}

type byAge []person
type byName []person

// I thinkit uses a combination of these methods to get the job done?
func (ba byAge) Len() int           { return len(ba) }
func (ba byAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba byAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

func (bn byName) Len() int           { return len(bn) }
func (bn byName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn byName) Less(i, j int) bool { return bn[i].First < bn[j].First }

// tbh this one is right over my head, lots of implicity work going on here
func customSort() {
	fmt.Println("------custom sorting------")
	p1 := person{"josh", "withers", 14}
	p2 := person{"tanya", "gladstone", 46}
	p3 := person{"mike", "ditka", 61}
	p4 := person{"amy", "pinkman", 22}

	people := []person{p1, p2, p3, p4}

	fmt.Println(people)
	sort.Sort(byAge(people)) // sorts by age, the integer (could probably substitute this)
	fmt.Println(people)
	sort.Sort(byName(people)) // sorts by name, the string
	fmt.Println(people)
}

// storying password information
func bcryptfunction() {
	fmt.Println("------bcrypt password encryption------")
	// generate password
	pswrd := `password`
	bs, err := bcrypt.GenerateFromPassword([]byte(pswrd), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(pswrd)
	fmt.Println(bs)         // hashed password as slice of bytes
	fmt.Println(string(bs)) // hashed password as string

	// log in by validating this password

	pswrd1 := `password`
	err = bcrypt.CompareHashAndPassword(bs, []byte(pswrd1))
	if err != nil {
		fmt.Println("Unsuccessful login.")
		return
	}
	fmt.Println("Successful login")

}
