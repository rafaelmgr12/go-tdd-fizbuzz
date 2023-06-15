package main

import "testing"

func TestFizzBuzz(t *testing.T) {
	// arrange
	input := 1

	// act
	got := FizzBuzz(input)

	// assert
	if got != "1" {
		t.Errorf(`expected "1" but got %q`, got)
	}
}
