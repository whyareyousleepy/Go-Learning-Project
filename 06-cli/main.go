package main

import (
	"flag"
	"fmt"
)

var name = flag.String("name", "World", "Name to say hello to")
var inSpanish bool

func init() {
	flag.BoolVar(&inSpanish, "spanish", false, "Use spanish language")
	flag.BoolVar(&inSpanish, "s", false, "Use spanish language")
	flag.Parse()
}

func main() {
	if inSpanish {
		fmt.Printf("Hola %s!\n", *name)
	} else {
		fmt.Printf("Hello %s!\n", *name)
	}
}
