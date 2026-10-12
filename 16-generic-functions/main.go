package main

import (
	"fmt"
	"unicode"
)

func filter[T any](items []T, fx func(T) bool) []T {
	var res []T

	for _, v := range items {
		if fx(v) {
			res = append(res, v)
		}
	}

	return res
}

type Numeric interface {
	int8 | int16 | int32 | int64 | float32 | float64
}

func filterPositive[T Numeric](items []T) []T {
	var filtered []T

	for _, v := range items {
		if v > 0 {
			filtered = append(filtered, v)
		}
	}

	return filtered
}

func main() {

	// 1
	strings := []string{"Hello", "my", "name", "is", "Indigo", "Montoya"}

	strings = filter(strings, func(s string) bool {
		return unicode.IsUpper(rune(s[0]))
	})

	fmt.Println(strings)

	// 2
	ints := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	ints = filter(ints, func(i int) bool {
		return i%3 == 0
	})

	fmt.Println(ints)

	// 3
	numInts := []int8{1, 2, -1, 0, -10, 11, 12}
	numInts = filterPositive(numInts)
	fmt.Println(numInts)

	floats := []float32{-0.5, 0, 1.5, 123.45, -100.1}
	floats = filterPositive(floats)
	fmt.Println(floats)
}
