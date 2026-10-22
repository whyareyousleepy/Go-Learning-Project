package main

import (
	"log"
	"time"
)

func main() {
	done := make(chan bool)
	msg := make(chan string)

	go send(msg, done)

	timeout:= time.After(5 * time.Second)
	for {
		select {
		case m := <-msg:
			log.Println(m)
		case <-timeout:
			done <- true
			return
		}
	}
}

func send(ch chan<- string, done <-chan bool) {
	for {
		select {
		case <-done:
			log.Println("Done")
			close(ch)
			return
		default:
			ch <- "hello"
			time.Sleep(500 * time.Millisecond)
		}
	}
}
