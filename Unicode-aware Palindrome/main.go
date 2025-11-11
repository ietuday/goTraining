package main

import (
	"fmt"
	"unicode"
)

func IsPalindrome(s string) bool {
	// 1) Build a []rune with only letters/digits, lowercased
	buf := make([]rune, 0, len(s))
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			buf = append(buf, unicode.ToLower(r))
		}
	}
	// 2) Two-pointer check from both ends
	i, j := 0, len(buf)-1
	for i < j {
		if buf[i] != buf[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	tests := []struct {
		in   string
		want bool
	}{
		{"racecar", true},
		{"RaceCar", true},
		{"Åbba, ba!", true},
		{"नयन", true},
		{"hello", false},
		{"No 'x' in Nixon", true},
	}

	for _, t := range tests {
		got := IsPalindrome(t.in)
		fmt.Printf("%q -> %v (want %v)\n", t.in, got, t.want)
	}
}
