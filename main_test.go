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
			"ShouldReturnTwo_WhenTwoIsPassed",
			2,
			"2",
		},
		{
			"ShouldReturnFour_WhenFourIsPassed",
			4,
			"4",
		},
		{
			"ShouldReturnFizz_WhenThreeIsPassed",
			3,
			"Fizz",
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
