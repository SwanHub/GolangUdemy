package main

import (
	"fmt"
)

func main() {
	exercise1()  // composite literal array + ranging
	exercise2()  // compositie literal slice + ranging
	exercise3()  // slicing slices
	exercise4()  // appending values to slices
	exercise5()  // deleting valuse from slices using append
	exercise6()  // using make function + capacity, length with slices
	exercise7()  // slice of slices (matrixes)
	exercise8()  // map of slices
	exercise9()  // add values to map, print with range
	exercise10() // delete value from map directly by key value
}

func exercise1() {
	fmt.Println("--------exercise 1---------")
	a := [5]int{1, 2, 3, 4, 5}
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", a)
}

func exercise2() {
	fmt.Println("--------exercise 2---------")
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for _, v := range s {
		fmt.Println(v)
	}
	fmt.Printf("%T\n", s)
}

func exercise3() {
	fmt.Println("--------exercise 3---------")
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(s[:5])
	fmt.Println(s[5:])
	fmt.Println(s[2:7])
	fmt.Println(s[1:6])
}

func exercise4() {
	fmt.Println("--------exercise 4---------")
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	s = append(s, 52)
	fmt.Println(s)
	s = append(s, 53, 54, 55)
	fmt.Println(s)

	y := []int{56, 57, 58, 59, 60}
	s = append(s, y...)
	fmt.Println(s)
}

func exercise5() {
	fmt.Println("--------exercise 5---------")
	s := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(s[:3], s[6:]...)
	fmt.Println(y)
}

func exercise6() {
	fmt.Println("--------exercise 6---------")
	s := make([]string, 50, 102)
	s = append(s, `Alabama`, ` Alaska`, ` Arizona`, ` Arkansas`, ` California`, ` Colorado`, ` Connecticut`, ` Delaware`, ` Florida`, ` Georgia`, ` Hawaii`, ` Idaho`, ` Illinois`, ` Indiana`, ` Iowa`, ` Kansas`, ` Kentucky`, ` Louisiana`, ` Maine`, ` Maryland`, ` Massachusetts`, ` Michigan`, ` Minnesota`, ` Mississippi`, ` Missouri`, ` Montana`, ` Nebraska`, ` Nevada`, ` New Hampshire`, ` New Jersey`, ` New Mexico`, ` New York`, ` North Carolina`, ` North Dakota`, ` Ohio`, ` Oklahoma`, ` Oregon`, ` Pennsylvania`, ` Rhode Island`, ` South Carolina`, ` South Dakota`, ` Tennessee`, ` Texas`, ` Utah`, ` Vermont`, ` Virginia`, ` Washington`, ` West Virginia`, ` Wisconsin`, ` Wyoming`)
	for i := 0; i < len(s); i++ {
		fmt.Println(i, s[i])
	}
	fmt.Println(len(s))
	fmt.Println(cap(s))
}

func exercise7() {
	fmt.Println("--------exercise 7---------")
	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}
	combo := [][]string{a, b}
	fmt.Println(combo)
}

func exercise8() {
	fmt.Println("--------exercise 8---------")
	jb := []string{`Shaken, not stirred`, `Martinis`, `Dudes`}
	mp := []string{`James Bond`, `Literature`, `Computer Science`}
	nd := []string{`Being evil`, `Ice cream`, `Sunsets`}
	m := map[string][]string{
		`bond_james`:      jb,
		`moneypenny_miss`: mp,
		`no_dr`:           nd,
	}
	fmt.Println(m)
}

func exercise9() {
	fmt.Println("--------exercise 9---------")
	jb := []string{`Shaken, not stirred`, `Martinis`, `Dudes`}
	mp := []string{`James Bond`, `Literature`, `Computer Science`}
	nd := []string{`Being evil`, `Ice cream`, `Sunsets`}
	m := map[string][]string{
		`bond_james`:      jb,
		`moneypenny_miss`: mp,
		`no_dr`:           nd,
	}
	m[`hatter_mad`] = []string{`Skateboarding`, `bad jokes`, `foosball`}
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func exercise10() {
	fmt.Println("--------exercise 10---------")
	jb := []string{`Shaken, not stirred`, `Martinis`, `Dudes`}
	mp := []string{`James Bond`, `Literature`, `Computer Science`}
	nd := []string{`Being evil`, `Ice cream`, `Sunsets`}
	m := map[string][]string{
		`bond_james`:      jb,
		`moneypenny_miss`: mp,
		`no_dr`:           nd,
	}
	delete(m, `bond_james`)
	for k, v := range m {
		fmt.Printf("Value for: %v\n", k)
		for _, v2 := range v {
			fmt.Printf("\t%v\n", v2)
		}
	}
}
