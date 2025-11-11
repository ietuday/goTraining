package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	fmt.Println("Day7 Started")
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	result, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 0 =", result)
	}

	fmt.Println("\nPanic and Recover Example")

	safeDivide := func (a, b int)  {
		defer func ()  {
			if r:= recover(); r !=nil {
				fmt.Println("Recovered from Panic: ",r )
			}
		}()
		if b==0{
			panic("division by zero caused panic")
		}
        fmt.Println("Result:", a/b)
	} 

	safeDivide(20, 4)
	safeDivide(20, 0)
	fmt.Println("Program still running after recover.")


}
