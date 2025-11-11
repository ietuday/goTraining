package main

import (
	"fmt"
	"sync"
	"time"
)

func printMessage(msg string, wg *sync.WaitGroup){
	defer wg.Done()

	for i:=1;i<=5; i++{
		fmt.Println(msg, i)
		time.Sleep(400 * time.Millisecond)
	}
}

func main(){
	var wg  sync.WaitGroup
	
	wg.Add(2)

	go printMessage("Hello", &wg)
	go printMessage("World", &wg)

	wg.Wait()
	fmt.Println("All Go Routines Completed! ")
}
