package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan bool)
	timeout := time.After(500 * time.Millisecond)

	go send(ch)

	for {
		select {
		case b, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed")
				return
			}
			fmt.Printf("Got message: %v\n", b)
		case <-timeout:
			fmt.Println("Timed out")
			return
		default:
			fmt.Println("yawn")
			time.Sleep(100 * time.Millisecond)
		}
	}

}

func send(ch chan bool) {
	time.Sleep(100 * time.Millisecond)
	ch <- true
	close(ch)
	fmt.Println("Sent and closed")
}
