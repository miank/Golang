// Implement worker pool

package main

import (
	"fmt"
	"sync"
	"time"
)

// This is the function where each go routine runs , it listens to the jobs channel.

func worker(id int, jobs <-chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for job := range jobs {
		fmt.Printf("Worker %d: processing jobs %d \n", id, job)
		time.Sleep(time.Millisecond * 500)
	}

}

func main() {

	numJobs := 5
	numWorkers := 2

	var wg sync.WaitGroup
	jobs := make(chan int)

	for w := 0; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, &wg)
	}

	// send jobs
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}

	close(jobs)

	wg.Wait()

	fmt.Println("All jobs done")

}
