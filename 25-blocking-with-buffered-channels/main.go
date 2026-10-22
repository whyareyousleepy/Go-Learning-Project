package main

import (
	"log"
	"time"
)

func main() {

	lock := make(chan bool, 1)
	for i := 1; i < 7; i++ {
		go worker(i, lock)
	}
	time.Sleep(2 * time.Second)
}

func worker(id int, lock chan bool) {
	log.Printf("%d wants the lock\n", id)
	lock <- true
	log.Printf("%d acquired lock\n", id)
	time.Sleep(200 * time.Millisecond)
	<-lock
	log.Printf("%d released lock\n", id)
}
