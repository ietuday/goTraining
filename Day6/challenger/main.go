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
	return 2 * (r.width + r.height)
}

func updateRectangle(rect *Rectangle, w, h float64){
	rect.width = w
	rect.height = h
}

func main(){
	fmt.Println("Day6 Challenge")
	rect := Rectangle{width: 5, height: 4}

	fmt.Println("Before Update: ")
	fmt.Println("Width: ", rect.width, "Height: ",rect.height)
	fmt.Println("Area: ", rect.Area())
	fmt.Println("Perimeter: ", rect.Perimeter())

	updateRectangle(&rect, 10, 6)
	fmt.Println("Update Update: ")
	fmt.Println("Width: ", rect.width, "Height: ",rect.height)
	fmt.Println("Area: ", rect.Area())
	fmt.Println("Perimeter: ", rect.Perimeter())

}
