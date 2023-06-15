package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz_One(t *testing.T) {
	// arrange
	input := 1

	// act
	got := FizzBuzz(input)

	// assert
	assert.Equal(t, strconv.Itoa(input), got)
}

func TestFizzBuzz_Two(t *testing.T) {
	// arrange
	input := 2

	// act
	got := FizzBuzz(input)

	// assert
	assert.Equal(t, strconv.Itoa(input), got)
}

func TestFizzBuzz_Four(t *testing.T) {
	// arrange
	input := 4

	// act
	got := FizzBuzz(input)

	// assert
	if got != "4" {
		t.Errorf(`expected "4" but got %q`, got)
	}
}
