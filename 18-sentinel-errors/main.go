package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

const MAX_TIMEOUTS = 5

var TimeoutErr = errors.New("Timeout error")
var RejectedErr = errors.New("Rejected error")

func main() {
	response, err := SendRequest("Hello")

	if errors.Is(err, TimeoutErr) {
		timeouts := 0
		for errors.Is(err, TimeoutErr) {
			timeouts++
			fmt.Println("Timeout. Retrying")
			if timeouts == MAX_TIMEOUTS {
				panic("too many timeouts!")
			}
			response, err = SendRequest("Hello")
		}
	}

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(response)
	}
}

func SendRequest(req string) (string, error) {

	switch rand.IntN(3) % 3 {
	case 0:
		return "Success", nil
	case 1:
		return "", RejectedErr
	default:
		return "", TimeoutErr
	}

}
