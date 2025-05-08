// Task:
// Write a Go program that spawns two goroutines. One goroutine should calculate the sum of integers from 1 to 100,
// and the other should calculate the product of integers from 1 to 100. Both results should be sent to a channel,
// and the main function should print both results once they are received.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	ch := make(chan interface{}, 2)

	go func() {
		// Calculate sum
		sum := 0
		for i := 1; i <= 100; i++ {
			sum += i
		}
		ch <- sum

		// Calculate product
		prod := big.NewInt(1)
		for i := 1; i <= 100; i++ {
			prod.Mul(prod, big.NewInt(int64(i)))
		}
		ch <- prod
	}()

	// Receive and print both results
	for i := 0; i < 2; i++ {
		result := <-ch
		switch v := result.(type) {
		case int:
			fmt.Println("Sum:", v)
		case *big.Int:
			fmt.Println("Product:", v)
		}
	}

}
