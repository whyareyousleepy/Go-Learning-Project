package main

import (
	"bytes"
	"net/http"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			http.Post(
				"http://localhost:8080/users",
				"application/json",
				bytes.NewBufferString(`{"name":"test"}`),
			)
		}()
	}

	wg.Wait()
}