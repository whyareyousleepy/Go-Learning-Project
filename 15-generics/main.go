package main

import "fmt"

type Cat struct {
	wearsBow bool
}

type Dog struct {
	canFetch bool
}

type AnimalType interface {
	Cat | Dog
}

type Animal[T AnimalType] struct {
	value       T
	AnimalNoise func() string
}

func (a Animal[T]) speak() {
	fmt.Printf("we got a %T\n", a.value)
	fmt.Println(a.AnimalNoise())
}

func main() {
	catAnimal := Animal[Cat]{
		value: Cat{},
		AnimalNoise: func() string {
			return "meow"
		},
	}

	dogAnimal := Animal[Dog]{
		value: Dog{},
		AnimalNoise: func() string {
			return "woof"
		},
	}

	catAnimal.speak()
	dogAnimal.speak()
}
