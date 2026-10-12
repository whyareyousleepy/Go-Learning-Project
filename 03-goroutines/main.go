package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var ch chan int

func main() {
	ch = make(chan int, 2)
	wg.Add(2)
	go getValue()
	go sendValue(5)
	wg.Wait()

	fmt.Println("Main finished!")
}

func sendValue(n int) {
	defer wg.Done()
	ch <- n
	ch <- n + 1
	close(ch)
}

func getValue() {
	defer wg.Done()
	for {
		value, ok := <-ch
		if !ok {
			break
		}
		fmt.Printf("Got value: %d\n", value)
	}
}
