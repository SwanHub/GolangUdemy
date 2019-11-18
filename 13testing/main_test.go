package main

import (
	"testing"
)

//two forms of table testing. Two forms of structs.

func TestSum(t *testing.T) {

	type test struct {
		data   []int
		answer int
	}

	tests := []test{
		test{[]int{2, 3}, 5},
		test{[]int{4, 5}, 9},
		test{[]int{6, 7}, 13},
	}

	for _, v := range tests {
		x := Sum(v.data...)
		if x != v.answer {
			t.Errorf("Expected: %v | Got: %v\n", v.answer, x)
		}
	}

}

func TestModulo(t *testing.T) {
	var tests = []struct {
		input1   int
		input2   int
		expected int
	}{
		{10, 4, 2},
		{137, 8, 1},
		{200, 2, 0},
	}

	for _, test := range tests {
		if output := Modulo(test.input1, test.input2); output != test.expected {
			t.Errorf("Test Failed: Expected %v mod %v to equal %v, but got %v", test.input1, test.input2, test.expected, output)
		}
	}
}
