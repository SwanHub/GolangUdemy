package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	got := Add(2)
	want := 4
	if got != want {
		t.Error("wrong number!!!!")
	}
}
