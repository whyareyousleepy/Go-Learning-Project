package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type SmallInt int8

func doubler[T constraints.Integer](value T) T {
	return value * T(2)
}

func main() {
	var four = SmallInt(4)

	fmt.Println(doubler(four))
}
