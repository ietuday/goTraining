package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.name, p.age)
}

func (p *Person) haveBirthday() {
	p.age++
	fmt.Printf("Happy Birthday %s! You are now %d.\n", p.name, p.age)
}

func main() {
	fmt.Println("Day5 started")
	person1 := Person{name: "Udayaditya Singh", age: 33}
	person1.greet()
	person1.haveBirthday()
}
