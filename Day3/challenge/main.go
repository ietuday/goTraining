package main

import "fmt"

func multiplyAndAdd(a, b int) (int, int) {
	product := a * b
	sum := a + b
	return product, sum
}

func main() {
	product, sum := multiplyAndAdd(7, 3)
	fmt.Println("Product:", product)
	fmt.Println("Sum:", sum)
}
