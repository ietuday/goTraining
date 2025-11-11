package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	sum := 0

	for _, num := range numbers {
		sum += num
	}

	fmt.Println("Numbers: ", numbers)
	fmt.Println("Sum of numbers", sum)

	capitals := map[string]string{
		"India": "New Delhi",
		"USA":   "Washington D.C",
		"Japan": "Tokyo",
	}
	fmt.Println("Country Capital List")
	for country, capcapital := range capitals {
		fmt.Printf("%s => %s\n", country, capcapital)
	}
}
