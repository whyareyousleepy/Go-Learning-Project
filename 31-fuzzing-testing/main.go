package main

import (
	"fmt"
	"strings"
)

func summedRunceCodes(input string) int16 {
	value := 0
	inRunes := strings.Map(func(r rune) rune {
		return r
	}, input)

	for r := range inRunes {
		value += int(inRunes[r])
	}

	return int16(value)
}

func main() {
	testString := "i am test string"
	output := summedRunceCodes(testString)
	fmt.Println(output)
}
