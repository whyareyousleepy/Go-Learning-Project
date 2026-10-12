package main

import (
	"fmt"
	"reflect"
)

type Animal struct {
	name string `help:"the name of the animal must be cat or dog"`
}

func (a Animal) speak() string {
	switch a.name {
	case "cat":
		return "meow"
	case "dog":
		return "woof"
	default:
		if member, ok := reflect.TypeOf(a).FieldByName("name"); ok {
			return fmt.Sprintf("Invalid name: %s", member.Tag.Get("help"))
		}
		return "nondescript animal noise"
	}
}

func main() {
	a := Animal{
		name: "cat1",
	}

	fmt.Println(a.speak())
}
