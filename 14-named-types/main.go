package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

type Shuffleable interface {
	shuffle()
	contents() string
}

type shuffleString string

func NewShuffleString(init string) *shuffleString {
	var s = shuffleString(init)
	return &s
}

func (s *shuffleString) shuffle() {
	tmp := strings.Split(string(*s), "")
	rand.Shuffle(len(tmp), func(i, j int) {
		tmp[i], tmp[j] = tmp[j], tmp[i]
	})
	*s = shuffleString(strings.Join(tmp, ""))
}

func (s *shuffleString) contents() string {
	return string(*s)
}

type shuffleSlice []interface{}

func (sl shuffleSlice) contents() string {
	data, _ := json.Marshal(sl)
	return fmt.Sprintf("%v", string(data))
}

func (sl shuffleSlice) shuffle() {
	rand.Shuffle(len(sl), func(i, j int) {
		sl[i], sl[j] = sl[j], sl[i]
	})
}

func main() {
	var s Shuffleable
	s = NewShuffleString("my name is Test")
	s.shuffle()
	fmt.Println(s.contents())

	s = shuffleSlice{1, 2, 3, 4, 5}
	s.shuffle()
	fmt.Println(s.contents())
}
