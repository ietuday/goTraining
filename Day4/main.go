package main

import "fmt"

func main(){
	var numbers [5]int
	numbers[0] = 10
	numbers[1] = 20
	numbers[2] = 30

	fmt.Println("Array: ", numbers)
	fmt.Println("First element: ", numbers[0])

	fruits :=[] string {"Apple", "Banana", "Mango"}
	fruits = append(fruits, "Orange")
	fmt.Println("Slice: ", fruits)
	fmt.Println("Length: ", len(fruits))

	for i, fruit := range fruits {
		fmt.Printf("Fruit %d: %s\n", i, fruit)
	}

	ages :=map [string] int {
		"Udayaditya": 33,
		"Niketa": 30,
	}

	ages ["Riya"] = 25

	fmt.Println("Ages: ", ages)
	fmt.Println("Age of Niketa: ", ages["Niketa"])

	for name, age := range ages {
		fmt.Printf("%s is %d year old\n", name, age)
	}

}



