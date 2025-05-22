// Implement a Worker Pool
// Problem Statement:

// Write a Go program that implements a worker pool with the following behavior:

// You have N worker goroutines (e.g., 3 workers).

// You have a list of tasks (for example, numbers from 1 to 10).

// Each worker takes a task from a job queue (channel), processes it (simulated by sleeping for some time), then prints the result.

// The main function should wait until all tasks are processed before exiting.

package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(worker int, jobs chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for x := range jobs {
		fmt.Printf("Worker %d processed job %d\n", worker, x)
		time.Sleep(2 * time.Millisecond)
	}
}

func main() {

	var wg sync.WaitGroup
	jobs := make(chan int, 3)
	work := 3

	wg.Add(work)

	for i := 1; i <= work; i++ {
		go Worker(i, jobs, &wg)
	}

	for i := 1; i <= 3; i++ {
		jobs <- i
	}
	close(jobs)

	wg.Wait()

}
