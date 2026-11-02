package main

import (
	"fmt"
	"strings"
)

func main() {

}

func fizzbuzz(n int) string {

	var fizzbuzzes []string

	for i := 1; i <= n; i++ {

		isThree := i%3 == 0
		isFive := i%5 == 0

		if isThree && isFive {
			fizzbuzzes = append(fizzbuzzes, "Fizzbuzz")
		} else if isThree {
			fizzbuzzes = append(fizzbuzzes, "Fizz")
		} else if isFive {
			fizzbuzzes = append(fizzbuzzes, "Buzz")
		} else {
			fizzbuzzes = append(fizzbuzzes, fmt.Sprintf("%d", i))
		}
	}

	return strings.Join(fizzbuzzes, ", ")

}
