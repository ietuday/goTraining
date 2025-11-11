package main

import (
	"fmt"
	"time"
)

func printNumber() {
	for i := 0; i < 5; i++ {
		fmt.Println("Numbers: ", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func printLettes() {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Printf("Letters: %c \n", ch)
		time.Sleep((500 * time.Millisecond))
	}
}

func main() {
	fmt.Println("Day15 started")
	fmt.Println("Start")

	go printNumber()
	go printLettes()

	time.Sleep(3 * time.Second)

	fmt.Println("End")

}
