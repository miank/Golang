// **Task**:
// Write a Go program that uses a **`sync.Mutex`** to safely update a shared counter across 10 goroutines.
// Each goroutine should increment the counter 100 times.
// Use a **`sync.WaitGroup`** to ensure the main goroutine waits for all goroutines to finish before printing the final counter value.

package main

import (
	"fmt"
	"sync"
)

var (
	counter int
	rw      sync.Mutex
	wg      sync.WaitGroup
)

func increment(r *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		r.Lock()
		counter++
		r.Unlock()
	}
}

func main() {

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go increment(&rw, &wg)
	}

	wg.Wait()
	fmt.Println("Counter Value :", counter)
}
