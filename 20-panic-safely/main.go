package main

import (
	"errors"
	"github.com/MasterMinds/cookoo/safely"
	"time"
)

func message() {
	println("Inside goroutine")
	panic(errors.New("This is a panic"))
}

func main() {
	safely.Go(message)
	println("Outside goroutine")
	time.Sleep(2 * time.Second)
}
