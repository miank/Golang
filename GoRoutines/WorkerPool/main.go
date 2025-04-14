// Worker Pool

// Implement a worker pool where multiple goroutines (workers) pick up tasks from a shared task queue (channel).

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Job struct {
	ID        int
	RandomNum int
}

type Result struct {
	JobID  int
	Square int
}

func worker(id int, jobs <-chan Job, result chan<- Result, wg *sync.WaitGroup) {

	defer wg.Done()
	for job := range jobs {
		fmt.Printf("Worker %d processing jobs %d \n", id, job.ID)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(500)))
		result <- Result{
			JobID: job.ID, Square: job.RandomNum * job.RandomNum}
	}
}

func main() {

	const (
		numJobs    = 10
		numWorkers = 3
	)

	jobs := make(chan Job, numJobs)
	results := make(chan Result, numJobs)

	var wg sync.WaitGroup

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, results, &wg)
	}

	// send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- Job{ID: j, RandomNum: rand.Intn(100)}
	}

	// collect result
	for result := range results {
		fmt.Printf("Result: JobID=%d, Square=%d\n", result.JobID, result.Square)
	}

}
