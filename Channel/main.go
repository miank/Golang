package main

import (
	"fmt"
	"log"
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

}
