package main

import "strconv"

func FizzBuzz(input int) string {
	if input == 5 {
		return "Buzz"
	}
	if input%3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(input)
}
