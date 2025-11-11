package main

import "fmt"

func main() {
	fmt.Println("Start of main")

	defer fmt.Println("This runs at the end (1)")
	defer fmt.Println("This runs at the end (2)")

	fmt.Println("Doing some work...")

	fmt.Println("End of main function")
}
