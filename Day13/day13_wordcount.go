package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Define a flag for filename
	filename := flag.String("file", "", "Path to the text file to analyze")
	flag.Parse()

	// Check if file provided
	if *filename == "" {
		fmt.Println("Usage: go run day13_wordcount.go -file=<filename>")
		return
	}

	// Open the file
	file, err := os.Open(*filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Initialize counters
	var lines, words, chars int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines++
		chars += len(line)
		words += len(strings.Fields(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("📄 File: %s\n", *filename)
	fmt.Printf("➡ Lines: %d\n", lines)
	fmt.Printf("➡ Words: %d\n", words)
	fmt.Printf("➡ Characters: %d\n", chars)
}
