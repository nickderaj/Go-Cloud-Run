package main

import (
	"testing"
)

// Add - Dummy Fn for testing, remove this later
func Add(x, y int) (res int) {
	return x + y
}

func TestHelloWorld(t *testing.T) {
	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
