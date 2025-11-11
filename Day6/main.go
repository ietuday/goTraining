package main

import "fmt"

func increment(num *int) {
	*num = *num + 1
}

type Person struct {
	name string 
	age int
}

func (p *Person) birthday(){
	p.age ++
	fmt.Printf("Happy Bithday %s! You are now %d years old\n", p.name, p.age)
}

func main(){
	fmt.Println("Day6 Started")

	x:=10
	fmt.Println("Before Increment: ", x)
	increment(&x)
	fmt.Println("After Increment: ", x)

	person := Person {name: "Udayaditya", age: 33}

	person.birthday()
	person.birthday()
}