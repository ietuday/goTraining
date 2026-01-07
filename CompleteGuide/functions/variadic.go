package main

import "fmt"



func sumup(num ...int) int{
	sum:=0

	for index, val := range num {
		fmt.Println(index)
		sum+=val
		
	}
	return sum
}