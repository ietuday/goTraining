package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Hello, my name is Udayaditya Singh")

	today := time.Now()

	fmt.Println("Today date is : ", today.Format("02-Jan-2006"))

}
