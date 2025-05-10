// Implement a Go program where you have **3 producers** (goroutines) each sending 5 integers into a **buffered channel**.
// The main goroutine will receive and print the integers until the channel is empty.
// Ensure that the main goroutine properly waits for all producers to finish before closing the channel.

package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan int, 15)
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		for i := 1; i <= 5; i++ {
			ch <- i
		}

	}()

	go func() {
		defer wg.Done()
		for i := 6; i <= 10; i++ {
			ch <- i
		}
	}()

	go func() {
		defer wg.Done()
		for i := 10; i <= 15; i++ {
			ch <- i
		}
	}()

	// Important Remember

	go func() {
		wg.Wait()
		close(ch)
	}()

	// Another way for waiting for the signal to be done.

	// Wait for all 3 producers to signal done
	// go func() {
	// 	for i := 0; i < 3; i++ {
	// 		<-done
	// 	}
	// 	close(ch)
	// }()

	for val := range ch {
		fmt.Println("Received:", val)
	}

}
