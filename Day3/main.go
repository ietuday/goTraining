package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func divide(a int, b int) (int, int) {
	quotient := a / b
	remainder := a % b
	return quotient, remainder

}
func main() {
	fmt.Println("Day3 with GO")
	sum := add(4, 5)
	fmt.Println("Sum: ", sum)

	q, r := divide(20, 3)

	fmt.Println("Quotient: ", q, "remainder: ", r)

	qOnly, _ := divide(22, 5)
	fmt.Println("Only Quotient:", qOnly)

}
