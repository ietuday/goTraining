package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func speak(word string, times int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= times; i++ {
		// jitter to make interleaving obvious
		time.Sleep(time.Duration(100+rand.Intn(400)) * time.Millisecond)
		fmt.Printf("%s  [%d/5]  %s\n",
			time.Now().Format("15:04:05.000"), i, word)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	wg.Add(3)

	go speak("Go", 5, &wg)
	go speak("is", 5, &wg)
	go speak("awesome", 5, &wg)

	wg.Wait()
	fmt.Println("✅ All goroutines completed.")
}
