// Print Numbers Using Goroutines in Order
// Problem Statement:

// Write a Go program that uses two goroutines to print even and odd numbers from 1 to 10 in the correct order.

package main

import (
	"fmt"
	"sync"
)

func Even(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 10; i++ {
		<-ch
		if i%2 == 0 {
			fmt.Println(i)
		}
		ch <- 1
	}

}

func Odd(ch chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
		ch <- 1
		<-ch
	}

}

func main() {

	ch := make(chan int)

	var wg sync.WaitGroup
	wg.Add(2)

	go Even(ch, &wg)
	go Odd(ch, &wg)

	wg.Wait()
	close(ch)

}
