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
