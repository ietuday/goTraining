package main

import (
	"fmt"
	"unicode"
)

func isPalidrome(s string) bool {
	var cleaned []rune
	for _, r := range s {
		if unicode.IsSpace(r) {
			continue
		}
		cleaned = append(cleaned, unicode.ToLower(r))
	}

	i, j := 0, len(cleaned)-1
	for i < j {
		if cleaned[i] != cleaned[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	tests := []string{
		"Race car",
		"Never odd or even",
		"hello",
		"A man a plan a canal Panama",
	}
	fmt.Println(tests)

	for _, t := range tests {
		fmt.Println(t)
		fmt.Printf("%q-> %v", t, isPalidrome(t))
		fmt.Println()
	}
}
