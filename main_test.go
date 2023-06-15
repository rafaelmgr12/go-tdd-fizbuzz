package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	// arrange
	input := 1

	// act
	got := FizzBuzz(input)

	// assert
	assert.Equal(t, strconv.Itoa(input), got)
}
