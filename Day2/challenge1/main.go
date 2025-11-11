package main

import "fmt"

func main() {
	number := 7
	const MAX = 10

	fmt.Printf("Multiplication Table of %d\n", number)
	for i := 1; i <= MAX; i++ {
		fmt.Printf("%d x %d = %d\n", number, i, number*i)
	}
}
