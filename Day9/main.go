package main

import (
	"fmt"

	"example.com/day9/mathutil"
)

func main() {
	fmt.Println("Day9 started")
	greet("Uday")
	a, b := 7, 3
	fmt.Println("Add: ", mathutil.Add(a, b))
	fmt.Println("Mul: ", mathutil.Mul(a, b))
}
