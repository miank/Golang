package main

import (
	"fmt"
	"sync"
	"time"
)

// If you must use default, only do it for non-blocking checks.
// Otherwise, it’ll skip even if channels aren’t ready.

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan string)
	ch2 := make(chan string)

	wg.Add(2)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		ch1 <- "Hello from Channel1"
	}()

	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		ch2 <- "Hello from Channel2"
	}()

	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}
	wg.Wait()

}
