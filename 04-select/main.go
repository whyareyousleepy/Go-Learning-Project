package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var ch chan int

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go send(ch1, 1)
	go send(ch2, 2)

	select {
	case message := <-ch1:
		fmt.Println("Ch1: ", message)
	case message := <-ch2:
		fmt.Println("Ch2: ", message)
	}

	select {
	case message := <-ch1:
		fmt.Println("Ch1: ", message)
	case message := <-ch2:
		fmt.Println("Ch2: ", message)
	}
}

func send(ch chan string, seconds time.Duration) {
	time.Sleep(seconds * time.Second)
	ch <- "Hello from channel!"
}
