package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	fmt.Println(ch)
	go send(ch)
	//	fmt.Println(<-ch)
	go receive(ch)
	time.Sleep(1000 * time.Millisecond)

}

func send(ch chan int) {
	ch <- 6
}

func receive(ch chan int) {
	fmt.Println("Data Received", <-ch)
}
