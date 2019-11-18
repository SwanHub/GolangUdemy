package main

import (
	"fmt"
	"testing"

	"github.com/SwanHub/Udemy/Golang-ToddMcleod/z-exercises/013-testing/dog"
)

func TestYears(t *testing.T) {
	d := dog.Years(4)
	if d != 28 {
		t.Errorf("you missed the mark on that one... completely wrong")
	}
}

func ExampleYears() {
	fmt.Println(dog.Years(4))
	// Output:
	// 28
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		dog.Years(4)
	}
}
