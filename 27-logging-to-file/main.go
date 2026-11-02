package main

import (
	"log"
	"os"
)

func main() {

	file, err := os.OpenFile("logs.txt", os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		panic("Could not open file")
	}

	log.SetOutput(file)
	log.SetFlags(log.LstdFlags|log.Lshortfile)
	log.Println("Display in UTC and use a short filename")

}
