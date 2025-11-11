package main

import "fmt"

type Rectangle struct {
	width float64
	height float64
}

func (r Rectangle) Area() float64{
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.width + r.height)
}

func main(){
	fmt.Println("CHallenge Day 5")
	rect := Rectangle{width: 5, height: 3}

	fmt.Println("Rectangle: ")
	fmt.Println("Width: ", rect.width, "Height: ", rect.height)
	fmt.Println("Area: ", rect.Area())
	fmt.Println("Perimeter: ", rect.Perimeter())
}
