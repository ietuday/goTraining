package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for index, num := range nums {
		fmt.Println("Index: ", index)
		fmt.Println("Number: ", num)

		if num%2 == 0 {
			fmt.Printf("%d is even", num)
		} else {
			fmt.Printf("%d is odd", num)
		}
		fmt.Println("")
	}
}
