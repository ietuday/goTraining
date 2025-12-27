package main

import (
	"bufio"
	"fmt"
	"strings"
)

func mustGetUserInput(reader *bufio.Reader, prompt string) string {
	for {
		value, err := getUserInput(reader, prompt)
		if err != nil {
			fmt.Println("❌", err)
			continue
		}
		return value
	}
}

func getUserInput(reader *bufio.Reader, prompt string) (string, error) {
	fmt.Print(prompt)

	text, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input")
	}

	text = strings.TrimSpace(text)
	if text == "" {
		return "", fmt.Errorf("input cannot be empty")
	}

	return text, nil
}
