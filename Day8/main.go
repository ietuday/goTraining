package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Reactangle struct {
	width, height float64
}

func (r Reactangle) Area() float64 {
	return r.width * r.height
}

func (r Reactangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return 3.14159 * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.radius
}

func printShapeDetails(s Shape){
	fmt.Printf("Area: %.2f, Perimeter: %.2f \n", s.Area(), s.Perimeter())
}

func main(){
	rect := Reactangle{width: 5, height: 4}
	circle:= Circle{radius: 5}

	fmt.Println("Reactangle Details: ")
	printShapeDetails(rect)

	fmt.Println("Circle Details: ")
	printShapeDetails(circle)

}