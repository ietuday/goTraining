package main

import "fmt"

func main() {
	var name string = "Udayaditya"
	age := 33

	fmt.Println("Name:", name)
	fmt.Println("Age:", age)

	const pi = 3.14159
	fmt.Println("Value of PI: ", pi)

	fmt.Println("Counting from 1 to 5")
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}

	fmt.Println("Countdown from 5 to 1")
	j := 5
	for j > 0 {
		fmt.Println(j)
		j--
	}

	fmt.Println("Breaking out of Loop:")
	k:=1
	for{
		if k >3 {
			break
		}
		fmt.Println("Loop Iteration", k)
		k++
	}

}
