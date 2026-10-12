package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name           string  `json:"animal_name"`
	ScientificName string  `json:"scientific_name"`
	Weight         float32 `json:"animal_weight"`
}

func main() {

	a := Animal{
		Name:           "cat",
		ScientificName: "Felis catus",
		Weight:         10.5,
	}

	output,err := json.Marshal(a)
	if err != nil {
		panic(fmt.Errorf("Error marshalling: ", a))
	}

	fmt.Println(string(output))

}