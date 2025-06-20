// Write a Go program that uses goroutines and channels to
// calculate the square of numbers from 1 to 10 concurrently. The program should:
// Launch a goroutine for each number to compute its square.
// Use a channel to collect the results.
// Print the results in the main function as they are received.

package main

import (
	"fmt"
	"sync"
)

func Square(num int, ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()
	ch <- num * num
}

func main() {

	ch := make(chan int, 10)
	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go Square(i, ch, &wg)
	}

	wg.Wait()
	close(ch)

	for val := range ch {
		fmt.Println(val)
	}

}
