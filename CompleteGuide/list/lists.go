package main

import "fmt"

func main() {
	hobbies := [3]string{"Sports", "Reading", "Movies"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	mainHobies :=hobbies[:2]
	fmt.Println(mainHobies)

	fmt.Println(cap(mainHobies))

	mainHobies = mainHobies[1:3]
	fmt.Println(mainHobies)

}
