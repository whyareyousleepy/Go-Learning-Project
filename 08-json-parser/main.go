package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type configuration struct {
	Enabled bool
	Path    string
}

func main() {

	file, err := os.Open("conf.json")
	if err != nil {
		fmt.Println("Error while opening file")
		return
	}
	defer file.Close()

	conf := configuration{}

	err = json.NewDecoder(file).Decode(&conf)
	if err != nil {
		fmt.Println("Error while reading file")
	}

	fmt.Println(conf.Path)
}
