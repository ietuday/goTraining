package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d processing job %d \n", id, job)
		result := job * job
		results <- result
	}
}

func main() {
	numsJob := 100
	numsWorker := 5

	jobs := make(chan int, numsJob)
	results := make(chan int, numsJob)

	var wg sync.WaitGroup

	for w := 1; w <= numsWorker; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	for j := 1; j <= numsJob; j++ {
		jobs <- j
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	for r := range results {
		fmt.Println("Result: ", r)
	}

}
