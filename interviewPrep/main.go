package main

import "fmt"

func main() {
	s := "Hello🙂"

	fmt.Println(len(s)) // bytes
	for _, r := range s {
		fmt.Println(string(r))
	}

	a := []int{1, 2, 3}
	b := a

	b = append(b, 4)

	fmt.Println(a, b)

}
