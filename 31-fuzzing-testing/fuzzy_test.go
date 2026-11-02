package main

import "testing"

func FuzzSummedRuneCodes(f *testing.F) {
	tests := []string{
		"i am trying things",
		"doing my best to find a way to break this!",
		"even adding emojis won't break it 😎",
	}

	for t := range tests {
		f.Add(tests[t])
	}

	f.Fuzz(func(t *testing.T, seed string) {
		got := summedRunceCodes(seed)
		if got < 0 {
			t.Errorf("how did this happen? somehow we got %d from string %s", got, seed)
		}
	})
}
