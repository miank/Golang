package main

import (
	"fmt"
	"time"
)

func main() {
	var c chan int
	fmt.Println(c)

	// To define a channel we have to use the inbuilt function make
	ch := make(chan int)
	fmt.Println(ch)

	// sending and receiving value from channel
	fmt.Println("Sending value to channel")
	go send(ch)

	fmt.Println("Receiving value from channel")
	go receive(ch)

	time.Sleep(time.Second * 1)

	// Sending and receing value in main
	ch1 := make(chan int)
	fmt.Println("Sending value to channel start")
	go send1(ch1)
	val := <-ch1
	fmt.Printf("Receiving value from channel finished. Value received: %d \n", val)

	// Buffered channels

	ch2 := make(chan int, 1)

	ch2 <- 2
	fmt.Println("Sending value to channel complete")

	val2 := <-ch2

	fmt.Printf("Receiving value from channel finished. Value received: %d\n", val2)

	// Receive channel is blocked when channel is empty

	ch3 := make(chan int, 1)
	ch3 <- 1
	fmt.Println("Sending to the channel is complete")
	fmt.Println("Send channel is blocked so rest of the code will not execute")
	val3 := <-ch3
	fmt.Printf("Receiving value from channel is complete %d", val3)
	//	val3 = <-ch3 // This will nlock the code

	fmt.Printf("Receving value from channel finished. Value received: %d\n", val3)
}

func send1(ch chan int) {
	ch <- 5
}

func send(ch chan int) {
	ch <- 5
}

func receive(ch chan int) {
	x := <-ch
	fmt.Println(x)
	fmt.Printf("Value received = %d in receive function\n", x)
}
