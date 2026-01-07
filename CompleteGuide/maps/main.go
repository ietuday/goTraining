package main

import "fmt"

func main() {
	username := make([]string, 2, 5)
	username = append(username, "Max")
	username = append(username, "Manuel")

	fmt.Println(username)

	courses := make(map[string]float64, 3)
	courses["go"] = 4.8
	fmt.Println(courses)

	for index, value := range username{
		fmt.Println(index, value)
	}

	for key, value := range courses{
		fmt.Println(key, value)
	}
}
