// Producer-Consumer with Buffered Channel
// Problem Statement:

// Write a Go program using goroutines and channels to simulate a Producer-Consumer scenario.

// One Producer goroutine should generate integers from 1 to 10 and send them into a buffered channel.

// One Consumer goroutine should read from the channel and print each number.

// Ensure that the producer finishes producing and the consumer finishes consuming without any deadlocks.

package main

import (
	"fmt"
	"sync"
)

func Producer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 10; i++ {
		ch <- i
	}
	close(ch)

}

func Consumer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for x := range ch {
		fmt.Println(x)

	}
}

func main() {

	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)
	go Producer(ch, &wg)
	go Consumer(ch, &wg)

	wg.Wait()
}
