package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		time.Sleep(150 * time.Millisecond) // simulate work
		results <- fmt.Sprintf("worker %d done job %d", id, j)
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan string)

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// feed jobs
	go func() {
		for j := 1; j <= 5; j++ {
			jobs <- j
		}
		close(jobs) // no more jobs
	}()

	// collect results asynchronously
	go func() {
		wg.Wait()
		close(results)
	}()

	for res := range results {
		fmt.Println(res)
	}
}
