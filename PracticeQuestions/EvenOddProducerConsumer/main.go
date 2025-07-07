// Go program that launches two goroutines:

// One sends even numbers to a channel

// Another sends odd numbers

// A single consumer reads and prints all numbers from the channel up to N (say N = 10)

package main

import (
	"fmt"
	"sync"
)

func Producer_1(N int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < N; i++ {
		if i%2 == 0 {
			fmt.Println("Producer ****", i)
			ch <- i
		}
	}
}

func Producer_2(N int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < N; i++ {
		if i%2 != 0 {
			fmt.Println("Producer ****", i)
			ch <- i
		}
	}
}

func Consumer(ch chan int) {
	for x := range ch {
		fmt.Println("Consumer ---->", x)
	}
}

func EvenOddProducerConsumer(N int) {

	ch := make(chan int, 10)
	var wg sync.WaitGroup
	wg.Add(2)

	go Producer_1(N, ch, &wg)
	go Producer_2(N, ch, &wg)
	go Consumer(ch)

	wg.Wait()
	close(ch)

}

func main() {
	EvenOddProducerConsumer(10)
}
