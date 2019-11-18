package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

func main() {
	exercise1() // Marshal data from struct to json string
	exercise2() // Unmarshal data from json string to sruct
	exercise3() // Marshal data from json string to sruct using Encoder and os.Stdout
	exercise4() // sorting slice of integers and strings
	exercise5() // sorting based on field in a struct
}

type user struct {
	First string
	Age   int
}

func exercise1() {
	fmt.Println("------exercise 1------")
	u1 := user{"James", 32}
	u2 := user{"Moneypenny", 27}
	users := []user{u1, u2}
	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(string(bs))
}

type verbosePerson struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

func exercise2() {
	fmt.Println("------exercise 2------")
	s := `[{"First":"James","Last":"Bond","Age":32,"Sayings":["Shaken, not stirred","Youth is no guarantee of innovation","In his majesty's royal service"]},{"First":"Miss","Last":"Moneypenny","Age":27,"Sayings":["James, it is soo good to see you","Would you like me to take care of that for you, James?","I would really prefer to be a secret agent myself."]},{"First":"M","Last":"Hmmmm","Age":54,"Sayings":["Oh, James. You didn't.","Dear God, what has James done now?","Can someone please tell me where James Bond is?"]}]`
	var verbosePeople []verbosePerson
	err := json.Unmarshal([]byte(s), &verbosePeople)
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println(verbosePeople)

	for _, v := range verbosePeople {
		fmt.Printf("My name is: %v %v\n", v.First, v.Last)
		fmt.Printf("I've been known to say:\n")
		for _, saying := range v.Sayings {
			fmt.Printf("\t%v\n", saying)
		}
	}
}

func exercise3() {
	fmt.Println("------exercise 3------")
	u1 := verbosePerson{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := verbosePerson{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	people := []verbosePerson{u1, u2}
	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error: ", err)
	}

	fmt.Println("****** method 1 for printing to terminal, uses string conv of slice of bytes")
	fmt.Println(string(bs))

	fmt.Println("****** method 2 for printing to terminal, uses encoding, quicker than the former")
	json.NewEncoder(os.Stdout).Encode(people)

}

func exercise4() {
	fmt.Println("------exercise 4------")
	xi := []int{5, 8, 2, 43, 17, 987, 14, 12, 21, 1, 4, 2, 3, 93, 13}
	xs := []string{"random", "rainbow", "delights", "in", "torpedo", "summers", "under", "gallantry", "fragmented", "moons", "across", "magenta"}

	fmt.Println(xi)
	sort.Ints(xi)
	fmt.Println(xi)

	fmt.Println(xs)
	sort.Strings(xs)
	fmt.Println(xs)

}

type byAge []verbosePerson
type byName []verbosePerson

// I thinkit uses a combination of these methods to get the job done?
func (ba byAge) Len() int           { return len(ba) }
func (ba byAge) Swap(i, j int)      { ba[i], ba[j] = ba[j], ba[i] }
func (ba byAge) Less(i, j int) bool { return ba[i].Age < ba[j].Age }

func (bn byName) Len() int           { return len(bn) }
func (bn byName) Swap(i, j int)      { bn[i], bn[j] = bn[j], bn[i] }
func (bn byName) Less(i, j int) bool { return bn[i].Last < bn[j].Last }

func exercise5() {
	fmt.Println("------exercise 5------")
	u1 := verbosePerson{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := verbosePerson{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	people := []verbosePerson{u1, u2}

	fmt.Println("---sorted by age")
	sort.Sort(byAge(people)) // sorts by age, the integer (could probably substitute this)
	for _, v := range people {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, s := range v.Sayings {
			fmt.Printf("\t%v\n", s)
		}
	}

	fmt.Println("---sorted by name")
	sort.Sort(byName(people)) // sorts by name, the string
	for _, v := range people {
		fmt.Println(v.First, v.Last, v.Age)
		sort.Strings(v.Sayings)
		for _, s := range v.Sayings {
			fmt.Printf("\t%v\n", s)
		}
	}

}
