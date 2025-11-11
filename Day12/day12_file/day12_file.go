package main

import (
	"bufio"
	"fmt"
	"os"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Create a file
	f, err := os.Create("example.txt")
	must(err)
	defer f.Close() // Always close at end

	writer := bufio.NewWriter(f)
	_, err = writer.WriteString("This is written safely using defer.\n")
	must(err)
	writer.Flush()
	fmt.Println("File written successfully!")

	// Read file back
	file, err := os.Open("example.txt")
	must(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Reading file content:")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	must(scanner.Err())
}
