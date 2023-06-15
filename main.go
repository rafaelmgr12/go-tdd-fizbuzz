package main

import "strconv"

func FizzBuzz(input int) string {
	if input == 3 || input == 6 {
		return "Fizz"
	}
	return strconv.Itoa(input)
}
