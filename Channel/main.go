package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	// Unbuffered channels
	ch2 := make(chan string, 1)
	// Buffered Channel
	ch6 := make(chan int, 16)

	// Send Data on a channel
	ch6 <- 2
	ch6 <- 4

	ch2 <- "HelloWorld"
	received := <-ch6
	fmt.Println("Data received from channel ", received)
	// channel closed
	close(ch6)

	x, ok := <-ch6
	if !ok {
		log.Println("Channel is empty or closed")
	}
	fmt.Println("The value of x is ", x)

	// For Loop channel
	ch := make(chan int, 6)
	ch <- 2
	ch <- 2
	ch <- 2
	ch <- 2
	ch <- 2
	ch <- 2
	close(ch)
	sum(ch)
	time.Sleep(time.Second * 1)

	// Receive and Fetch Value from a channel

	result := make(chan int, 1)

	go add(2, 4, result)

	value := result
	fmt.Printf("Total value : %d\n", <-value)
	close(result)
}

func add(a, b int, result chan int) {
	sumValue := a + b
	result <- sumValue
	return
}

func sum(ch chan int) {
	sum := 0
	for i := range ch {
		sum += i
	}
	fmt.Printf("Sum : %d \n", sum)
}
