// **Map with Goroutines**

// **Task**:
// Write a Go program where a **map** is populated concurrently by 3 goroutines.
// Each goroutine will update a key-value pair in the map, and the main goroutine should wait until all goroutines finish.
// Use proper synchronization (e.g., `sync.Mutex` or `sync.RWMutex`) to avoid race conditions.

package main

import (
	"fmt"
	"sync"
)

func main() {
	m1 := make(map[int]string)
	var mu sync.RWMutex
	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		defer wg.Done()
		mu.Lock()
		m1[1] = "Hello"
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		m1[2] = "World"
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		mu.Lock()
		m1[3] = "!!"
		mu.Unlock()
	}()

	wg.Wait()

	fmt.Println(m1[1], m1[2], m1[3])

}
