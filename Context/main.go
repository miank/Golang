package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Worker %d received cancellation signal. Exiting... \n ", id)
			return

		default:
			fmt.Printf("Worker %d is working ... \n ", id)
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	// Create a root context with a 2-second timeout

	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(ctx, &wg, 1)
	}

	time.Sleep(2 * time.Second)
	fmt.Println("Main: Cancelling the context")
	cancel()

	wg.Wait()
	fmt.Println("Main: All workers have exited")
}
