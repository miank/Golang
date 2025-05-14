package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs {
		fmt.Printf("Worker started %d job %d \n", id, job)
		time.Sleep(time.Second)
		fmt.Printf("Worker %d finished %d \n", id, job)
		results <- job * 2
	}
}

func main() {
	const numJobs = 5
	const numWorkers = 3

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs) // no of jobs that many results

	var wg sync.WaitGroup
	// First start the workers and the send then assign the jobs
	for j := 1; j <= numWorkers; j++ {
		wg.Add(1)
		go worker(j, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	wg.Wait()

	close(results)

	// Collect results
	for r := 1; r <= numJobs; r++ {
		fmt.Println("Result:", <-results)
	}

}
