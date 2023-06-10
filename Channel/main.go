package main

import (
	"fmt"
	"time"
)

func main() {
	var ch chan int
	fmt.Println(ch)

	c := make(chan int)
	fmt.Println(c)

	// sending and receiving value from channel
	fmt.Println("Sending value to channel")
	go send(ch)

	fmt.Println("Receiving value from channel")
	go receive(ch)

	time.Sleep(time.Second * 1)
}

func send(ch chan int) {
	ch <- 5
}

func receive(ch chan int) {
	x := <-ch
	fmt.Println(x)
	fmt.Printf("Value received = %d in receive function\n", x)
}
