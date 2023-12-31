package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFizzBuzz(t *testing.T) {
	// arrange
	testSuite := []struct {
		name     string
		input    int
		expected string
	}{
		{
			"ShouldReturnOne_WhenOneIsPassed",
			1,
			"1",
		},
		{
			"ShouldReturnFizz_WhenThreeIsPassed",
			3,
			"Fizz",
		},
		{
			"ShouldReturnBuzz_WhenFiveIsPassed",
			5,
			"Buzz",
		},
		{
			"ShouldReturnFizzBuzz_WhenFifteenIsPassed",
			15,
			"FizzBuzz",
		},
	}

	for _, tt := range testSuite {
		t.Run(tt.name, func(t *testing.T) {
			// act
			got := FizzBuzz(tt.input)

			// assert
			assert.Equal(t, tt.expected, got)
		})
	}
}
