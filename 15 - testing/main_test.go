package main

import "testing"

func TestAdd(t *testing.T) {
	// arrange
	l, r := 1, 2
	expect := 3

	// act
	got := Add(l, r)

	// assert
	if expect != got {
		t.Errorf("failed to add %v and %v. Expected %v and got %v", l, r, expect, got)
	}
}
