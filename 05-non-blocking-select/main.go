package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- 42
	}()

	select {
	case message := <-ch:
		fmt.Println("Ch1: ", message)
	default:
		fmt.Println("No value")
	}
}