package main

import (
	"fmt"
	"sync"
)

var val = 0

func worker(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	val = val + 1
	m.Unlock()

	// Wait Group will be done
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg, &m)
	}

	wg.Wait()
	fmt.Println("Value of x is ", val)
}
