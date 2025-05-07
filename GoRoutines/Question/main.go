// Task:
// Write a Go program that spawns two goroutines. One goroutine should calculate the sum of integers from 1 to 100,
// and the other should calculate the product of integers from 1 to 100. Both results should be sent to a channel,
// and the main function should print both results once they are received.

package main

import (
	"fmt"
	"math/big"
	"sync"
)

func task1(output chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}

	output <- sum

	close(output)

}

func task2(output1 chan *big.Int, wg *sync.WaitGroup) {
	defer wg.Done()

	prod := big.NewInt(1)
	for i := 1; i <= 100; i++ {
		prod.Mul(prod, big.NewInt(int64(i)))
	}
	output1 <- prod

	close(output1)
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan *big.Int)

	var wg sync.WaitGroup

	wg.Add(2)

	go task1(ch1, &wg)
	go task2(ch2, &wg)

	fmt.Println(<-ch1)
	fmt.Println(<-ch2)

	wg.Wait()

}
